# Runbook: Troubleshooting

## Pod stuck in Pending

```bash
kubectl describe pod <pod-name>
```

**Cause:** `pod has unbound immediate PersistentVolumeClaims`  
**Fix:** Ensure local-path-provisioner is installed and set as default storage class:
```bash
kubectl get storageclass
# Should show: local-path (default)
```

---

## API CrashLoopBackOff — DNS timeout

```bash
kubectl logs -l app.kubernetes.io/name=go-lab-api
# "hostname resolving error: lookup go-lab-db ... i/o timeout"
```

**Cause:** CoreDNS using `127.0.0.53` (systemd-resolved stub) as upstream — not reachable from pods.  
**Fix:** Patch CoreDNS ConfigMap to use the real upstream DNS:
```bash
UPSTREAM=$(ssh k8s@<VM_IP> 'resolvectl status | awk "/DNS Servers:/{print \$3; exit}"')
kubectl edit configmap coredns -n kube-system
# Change: forward . /etc/resolv.conf
# To:     forward . <UPSTREAM_IP>
kubectl rollout restart deployment coredns -n kube-system
```

---

## API CrashLoopBackOff — password authentication failed

```bash
kubectl logs -l app.kubernetes.io/name=go-lab-api
# "FATAL: password authentication failed for user"
```

**Cause:** `secret.dbPassword` in go-lab-api chart doesn't match the password set in go-lab-db.  
**Fix:**
```bash
helm upgrade go-lab-api helm/go_lab_api --set secret.dbPassword=<correct-password>
kubectl rollout restart deployment go-lab-api
```

---

## UI shows 502 Bad Gateway

**Cause:** go-lab-api service doesn't exist or API pod is not running.  
**Fix:**
```bash
kubectl get service go-lab-api
kubectl get pods -l app.kubernetes.io/name=go-lab-api
```

---

## SSH to VM hangs

**Cause:** ssh-agent not running.  
**Fix:**
```bash
eval $(ssh-agent -s) && ssh-add ~/.ssh/id_ed25519
```
