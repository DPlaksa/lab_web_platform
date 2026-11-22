# ADR-001: Use kubeadm for Kubernetes cluster provisioning

## Status

Accepted

## Context
<<<<<<< HEAD
We need a local Kubernetes cluster for development and learning purposes. Options considered:
- **minikube** — single-node, easy setup, but abstracts too much of the real Kubernetes internals
- **k3s** — lightweight, production-ready, but stripped-down control plane differs from full Kubernetes
- **kubeadm** — standard tool for bootstrapping full Kubernetes clusters, same as production environments
- **kind** — Kubernetes in Docker, good for CI but not representative of real infrastructure

## Decision
Use **kubeadm** to provision a full Kubernetes cluster inside a Multipass VM managed by Terraform.

## Consequences
- Full Kubernetes experience: real control plane, kubelet, kube-proxy, CNI
- Requires manual steps that managed Kubernetes (EKS, AKS) abstracts away — good for learning
- Heavier resource usage than k3s or minikube
- Direct path to reproducing the same setup on bare-metal or cloud VMs
=======

We need a local Kubernetes cluster for development and learning. The tool we pick should give real experience with how Kubernetes works in production, not abstract it away.

## Options considered

| Option | Notes |
|---|---|
| minikube | Easy setup, but hides too much of the internals |
| k3s | Lightweight and production-ready, but the control plane differs from full Kubernetes |
| **kubeadm** | Standard bootstrap tool, same as real production clusters |
| kind | Good for CI pipelines, not close enough to real infrastructure |

## Decision

Use kubeadm inside a Multipass VM managed by Terraform.

## Consequences

Good:
- Full Kubernetes: real control plane, kubelet, kube-proxy, CNI
- Nothing is abstracted away - good for learning
- Easy to reproduce on bare-metal or cloud VMs later

Bad:
- Uses more resources than k3s or minikube
- Requires more setup steps
>>>>>>> ce65d30 (docs: rewrite ADRs in MADR format, remove duplicate diagram title)
