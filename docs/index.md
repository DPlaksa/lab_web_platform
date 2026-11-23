# lab_web_platform

A production-style web platform built for learning and demonstrating modern DevOps practices.

## What it is

A three-component application (UI, API, and PostgreSQL) fully defined and configured declaratively as code, and deployable to a Kubernetes cluster using Terraform and Helm.

## Stack

| Layer | Technology |
|---|---|
| Infrastructure | Terraform + Multipass |
| OS | Ubuntu 26.04 |
| Kubernetes | kubeadm v1.31 |
| CNI | Flannel |
| Storage | local-path-provisioner |
| Package manager | Helm v3 |
| CI/CD | GitHub Actions |
| Registry | GitHub Container Registry (GHCR) |
| Language | Go (Fiber v2) |
| Database | PostgreSQL 16 |

## Repository structure

```
lab_web_platform/
├── apps/
│   ├── go_lab_api/       # Go API source + Dockerfile
│   ├── go_lab_ui/        # Static HTML/JS
│   └── go_lab_db/        # SQL migrations
├── helm/
│   ├── go_lab_api/       # Helm chart: API
│   ├── go_lab_ui/        # Helm chart: UI (nginx)
│   └── go-lab-db/        # Helm chart: PostgreSQL
├── infra/
│   └── k8s-node1/        # Terraform: Multipass VM + kubeadm
└── docs/                 # This documentation
```

## Quick start

See the [Deploy runbook](runbooks/deploy.md) for full instructions.

## Roadmap

- [ ] Terraform Cloud remote state
- [ ] Umbrella Helm chart
- [ ] Flux CD GitOps
- [ ] Flagger progressive deployments
- [ ] Terraform for Azure
- [ ] Prometheus + Grafana monitoring
- [ ] Ingress with TLS
