# ADR-003: Use independent Helm chart releases per component

## Status
Accepted

## Date
2026-11-22

## Context
The platform has three components (UI, API, DB) that evolve at different rates. Options considered:
- **Single Helm chart** — simple, but forces all components to release together
- **Umbrella chart** — parent chart with sub-charts, single install but still coupled versions
- **Independent charts with independent tags** — each component has its own release lifecycle

## Decision
Each component (`go-lab-ui`, `go-lab-api`, `go-lab-db`) has its own Helm chart and its own Git tag pattern (`ui-v*`, `api-v*`, `db-v*`), triggering independent GitHub Actions workflows.

## Consequences
- UI, API and DB can be released and rolled back independently
- GitHub Releases contains separate versioned artifacts per component
- Requires deploying components in order: DB → API → UI
- Umbrella chart planned as a future improvement to orchestrate the full stack deployment
