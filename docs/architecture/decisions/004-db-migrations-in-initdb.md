# ADR-004: Run DB migrations via PostgreSQL initdb mechanism

## Status

Accepted

## Context
<<<<<<< HEAD
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
=======

The database schema must exist before the API starts. The migration approach should keep concerns separated: the DB owns the schema, the API owns the business logic.

## Options considered

| Option | Notes |
|---|---|
| Migrations in app code | Simple, but the app creates its own schema - mixes concerns |
| Kubernetes Job | Clean separation, but adds complexity and ordering logic |
| **PostgreSQL initdb** | SQL files in `/docker-entrypoint-initdb.d` run on first boot automatically |

## Decision

Migration SQL is stored in a ConfigMap and mounted into `/docker-entrypoint-initdb.d`. The API checks on startup that the migration ran via `SELECT EXISTS` before accepting traffic.

## Consequences

Good:
- DB owns the schema, API owns the business logic
- Runs only once on first volume init, idempotent by design
- API fails fast with a clear error if migration has not run

Bad:
- Only runs on first init - adding future migrations needs a different approach (e.g. golang-migrate job)
>>>>>>> ce65d30 (docs: rewrite ADRs in MADR format, remove duplicate diagram title)
