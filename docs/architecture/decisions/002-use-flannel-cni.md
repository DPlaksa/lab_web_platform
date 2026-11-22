# ADR-002: Use Flannel as CNI plugin

## Status

Accepted

## Context
<<<<<<< HEAD
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
=======

Kubernetes requires a CNI plugin to give pods a network. For a single-node lab, the priority is simplicity over advanced features.

## Options considered

| Option | Notes |
|---|---|
| **Flannel** | Simple overlay network, works out of the box, no config needed |
| Calico | Supports NetworkPolicy and BGP routing, more complex to set up |
| Cilium | eBPF-based, most advanced, requires a newer kernel |

## Decision

Use Flannel with VXLAN backend and pod CIDR `10.244.0.0/16`.

## Consequences

Good:
- No configuration beyond `--pod-network-cidr` in kubeadm init
- Well documented and stable

Bad:
- No NetworkPolicy support
- Migration to Calico or Cilium needed if network policies are required later
>>>>>>> ce65d30 (docs: rewrite ADRs in MADR format, remove duplicate diagram title)
