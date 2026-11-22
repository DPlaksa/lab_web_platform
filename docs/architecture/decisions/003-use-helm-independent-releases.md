# ADR-003: Use independent Helm chart releases per component

## Status

Accepted

## Context
<<<<<<< HEAD
The platform has three components (UI, API, DB) that evolve at different rates. Options considered:
- **Single Helm chart** — simple, but forces all components to release together
- **Umbrella chart** — parent chart with sub-charts, single install but still coupled versions
- **Independent charts with independent tags** — each component has its own release lifecycle

## Decision
Each component (`go-lab-ui`, `go-lab-api`, `go-lab-db`) has its own Helm chart and its own Git tag pattern (`ui-v*`, `api-v*`, `db-v*`), triggering independent GitHub Actions workflows.
=======

The platform has three components (UI, API, DB) that change at different rates. The release strategy should let each component ship independently without forcing the others to release at the same time.

## Options considered

| Option | Notes |
|---|---|
| Single Helm chart | Simple, but all components must release together |
| Umbrella chart | One install command, but versions are still coupled |
| **Independent charts per component** | Each component has its own release cycle |

## Decision

Each component has its own Helm chart and its own Git tag pattern (`ui-v*`, `api-v*`, `db-v*`). Pushing a tag triggers only the GitHub Actions job for that component.
>>>>>>> ce65d30 (docs: rewrite ADRs in MADR format, remove duplicate diagram title)

## Consequences

Good:
- UI, API and DB can be released and rolled back independently
<<<<<<< HEAD
- GitHub Releases contains separate versioned artifacts per component
- Requires deploying components in order: DB → API → UI
- Umbrella chart planned as a future improvement to orchestrate the full stack deployment
=======
- Each GitHub Release contains only the artifacts for that component

Bad:
- Components must be deployed in order: DB first, then API, then UI
- Three separate `helm install` commands instead of one

Note: an umbrella chart is planned to wrap all three into a single install command.
>>>>>>> ce65d30 (docs: rewrite ADRs in MADR format, remove duplicate diagram title)
