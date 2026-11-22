# ADR-001: Use kubeadm for Kubernetes cluster provisioning

## Status
Accepted

## Date
2026-11-22

## Context
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
