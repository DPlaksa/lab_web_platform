# ADR-004: Run DB migrations via PostgreSQL initdb mechanism

## Status
Accepted

## Date
2026-11-22

## Context
Database schema needs to be initialized before the API starts. Options considered:
- **Migration in application code** — API creates tables on startup, simple but mixes concerns
- **Separate migration job** — Kubernetes Job runs before API, correct separation but adds complexity
- **PostgreSQL initdb** — SQL files mounted into `/docker-entrypoint-initdb.d` run automatically on first boot

## Decision
Migration SQL is stored in a ConfigMap and mounted into `/docker-entrypoint-initdb.d` in the PostgreSQL container. The API verifies on startup that the migration has been applied via `SELECT EXISTS`.

## Consequences
- Clean separation: DB owns schema, API owns business logic
- Migrations run only once on first volume initialization — idempotent by design
- Adding future migrations requires a separate strategy (e.g. migrate tool or init job)
- API fails fast with a clear error if migration hasn't run yet
