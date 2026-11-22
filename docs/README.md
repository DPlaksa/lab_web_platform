# Documentation

## Architecture

| Document | Description |
|---|---|
| [C4 Context](architecture/diagrams/c4-context.md) | System and external dependencies |
| [C4 Containers](architecture/diagrams/c4-containers.md) | Runtime components and communication |
| [Deployment](architecture/diagrams/deployment.md) | Full pipeline from code to cluster |

## Architecture Decision Records

| ADR | Decision |
|---|---|
| [ADR-001](architecture/decisions/001-use-kubeadm.md) | Use kubeadm for cluster provisioning |
| [ADR-002](architecture/decisions/002-use-flannel-cni.md) | Use Flannel as CNI plugin |
| [ADR-003](architecture/decisions/003-use-helm-independent-releases.md) | Independent Helm chart releases per component |
| [ADR-004](architecture/decisions/004-db-migrations-in-initdb.md) | Run DB migrations via PostgreSQL initdb |

## Runbooks

| Runbook | Description |
|---|---|
| [Deploy](runbooks/deploy.md) | Full deployment from scratch |
| [Troubleshooting](runbooks/troubleshooting.md) | Common issues and fixes |
