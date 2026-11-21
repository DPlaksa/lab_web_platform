package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type config struct {
	Port            string
	AppVersion      string
	DBHost          string
	DBPort          string
	DBName          string
	DBUser          string
	DBPassword      string
	DBSSLMode       string
	DBMaxOpenConns  int
	DBMaxIdleConns  int
	DBConnMaxLife   time.Duration
	ShutdownTimeout time.Duration
}

type pingResponse struct {
	Status     string `json:"status"`
	ClickCount int64  `json:"clickCount"`
	Version    string `json:"version"`
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok && val != "" {
		return val
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	val, ok := os.LookupEnv(key)
	if !ok || val == "" {
		return fallback
	}
	parsed, err := strconv.Atoi(val)
	if err != nil {
		log.Printf("warn: invalid int for %s=%q, using %d", key, val, fallback)
		return fallback
	}
	return parsed
}

func getEnvDuration(key string, fallback time.Duration) time.Duration {
	val, ok := os.LookupEnv(key)
	if !ok || val == "" {
		return fallback
	}
	parsed, err := time.ParseDuration(val)
	if err != nil {
		log.Printf("warn: invalid duration for %s=%q, using %s", key, val, fallback)
		return fallback
	}
	return parsed
}

func loadConfig() config {
	return config{
		Port:            getEnv("PORT", "8080"),
		AppVersion:      getEnv("APP_VERSION", "v1"),
		DBHost:          getEnv("DB_HOST", "localhost"),
		DBPort:          getEnv("DB_PORT", "5432"),
		DBName:          getEnv("DB_NAME", "lab_db"),
		DBUser:          getEnv("DB_USER", "lab_user"),
		DBPassword:      getEnv("DB_PASSWORD", "lab_password"),
		DBSSLMode:       getEnv("DB_SSLMODE", "disable"),
		DBMaxOpenConns:  getEnvInt("DB_MAX_OPEN_CONNS", 10),
		DBMaxIdleConns:  getEnvInt("DB_MAX_IDLE_CONNS", 5),
		DBConnMaxLife:   getEnvDuration("DB_CONN_MAX_LIFETIME", 30*time.Minute),
		ShutdownTimeout: getEnvDuration("SHUTDOWN_TIMEOUT", 10*time.Second),
	}
}

func (c config) dsn() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName, c.DBSSLMode,
	)
}

func connectAndVerifyDB(ctx context.Context, cfg config) (*sql.DB, error) {
	db, err := sql.Open("pgx", cfg.dsn())
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}

	db.SetMaxOpenConns(cfg.DBMaxOpenConns)
	db.SetMaxIdleConns(cfg.DBMaxIdleConns)
	db.SetConnMaxLifetime(cfg.DBConnMaxLife)

	// В K8s приложение и БД могут стартовать параллельно: делаем ретраи до готовности Postgres.
	var pingErr error
	for attempt := 1; attempt <= 15; attempt++ {
		pingCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
		pingErr = db.PingContext(pingCtx)
		cancel()
		if pingErr == nil {
			break
		}
		log.Printf("waiting for postgres (attempt %d/15): %v", attempt, pingErr)
		select {
		case <-ctx.Done():
			_ = db.Close()
			return nil, ctx.Err()
		case <-time.After(2 * time.Second):
		}
	}
	if pingErr != nil {
		_ = db.Close()
		return nil, fmt.Errorf("postgres ping failed: %w", pingErr)
	}

	// Убеждаемся, что миграция уже применена и нужная строка в button_stats существует.
	verifyCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var exists bool
	if err := db.QueryRowContext(verifyCtx, `SELECT EXISTS (SELECT 1 FROM button_stats WHERE id = 1)`).Scan(&exists); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("verify button_stats table: %w", err)
	}
	if !exists {
		_ = db.Close()
		return nil, fmt.Errorf("button_stats row with id=1 not found: run database migration first")
	}

	return db, nil
}

func main() {
	cfg := loadConfig()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	db, err := connectAndVerifyDB(ctx, cfg)
	if err != nil {
		log.Fatalf("db init error: %v", err)
	}
	defer db.Close()

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		ReadTimeout:           5 * time.Second,
		WriteTimeout:          5 * time.Second,
	})

	app.Get("/healthz", func(c *fiber.Ctx) error {
		reqCtx, cancel := context.WithTimeout(c.Context(), 2*time.Second)
		defer cancel()

		if err := db.PingContext(reqCtx); err != nil {
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
				"status": "unhealthy",
				"error":  "db unavailable",
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "ok",
			"version": cfg.AppVersion,
		})
	})

	app.Post("/api/ping", func(c *fiber.Ctx) error {
		reqCtx, cancel := context.WithTimeout(c.Context(), 3*time.Second)
		defer cancel()

		var count int64
		err := db.QueryRowContext(reqCtx, `
			UPDATE button_stats
			SET click_count = click_count + 1
			WHERE id = 1
			RETURNING click_count
		`).Scan(&count)
		if err != nil {
			log.Printf("error updating button_stats: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status": "error",
				"error":  "failed to increment counter",
			})
		}

		return c.Status(fiber.StatusOK).JSON(pingResponse{
			Status:     "ok",
			ClickCount: count,
			Version:    cfg.AppVersion,
		})
	})

	serverErrCh := make(chan error, 1)
	go func() {
		log.Printf("go_lab_api starting on :%s (version=%s, db=%s:%s/%s)",
			cfg.Port, cfg.AppVersion, cfg.DBHost, cfg.DBPort, cfg.DBName)
		if err := app.Listen(":" + cfg.Port); err != nil {
			serverErrCh <- err
		}
	}()

	select {
	case err := <-serverErrCh:
		log.Fatalf("server failed: %v", err)
	case <-ctx.Done():
		log.Printf("shutdown signal received, draining connections (timeout=%s)...", cfg.ShutdownTimeout)
		if err := app.ShutdownWithTimeout(cfg.ShutdownTimeout); err != nil {
			log.Printf("graceful shutdown error: %v", err)
		}
	}

	log.Println("go_lab_api stopped cleanly")
}
