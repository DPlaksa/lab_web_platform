# ADR-002: Use Flannel as CNI plugin

## Status
Accepted

## Date
2026-11-22

## Context
Kubernetes requires a CNI plugin to provide pod networking. Options considered:
- **Flannel** — simple overlay network, minimal configuration, widely used
- **Calico** — supports NetworkPolicy, BGP routing, more complex setup
- **Cilium** — eBPF-based, most advanced, requires newer kernel

## Decision
Use **Flannel** with VXLAN backend and pod CIDR `10.244.0.0/16`.

## Consequences
- Zero configuration required beyond `--pod-network-cidr` in kubeadm init
- No NetworkPolicy support — acceptable for a single-node lab environment
- If NetworkPolicy is needed in the future, migration to Calico or Cilium is straightforward
