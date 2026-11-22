# Runbook: Deploy lab_web_platform

## Prerequisites

- Terraform >= 1.17
- Multipass >= 1.17
- kubectl >= 1.31
- Helm >= 3.0
- SSH key added to GitHub (`~/.ssh/id_ed25519`)

## 1. Provision the VM and Kubernetes cluster

```bash
cd infra/k8s-node1
terraform init
terraform apply -auto-approve
```

Wait ~5 minutes for cloud-init to complete (kubeadm init, Flannel, local-path-provisioner).

Verify:
```bash
ssh k8s@<VM_IP> 'cloud-init status'
# Expected: status: done
```

## 2. Configure kubectl on your machine

```bash
scp k8s@<VM_IP>:/home/k8s/.kube/config ~/.kube/config-k8s-lab
export KUBECONFIG=~/.kube/config-k8s-lab
kubectl get nodes
# Expected: k8s-lab-n1   Ready   control-plane
```

## 3. Deploy the stack

Deploy in order — DB must be ready before API starts.

```bash
# 1. Database
helm install go-lab-db helm/go-lab-db
kubectl rollout status statefulset/go-lab-db

# 2. API
helm install go-lab-api helm/go_lab_api --set secret.dbPassword=<DB_PASSWORD>
kubectl rollout status deployment/go-lab-api

# 3. UI
helm install go-lab-ui helm/go_lab_ui
kubectl rollout status deployment/go-lab-ui
```

## 4. Verify

```bash
kubectl get pods
# All three pods should be Running

kubectl run curl-test --image=curlimages/curl --rm -it --restart=Never -- \
  curl -s http://go-lab-api:8080/healthz
# Expected: {"status":"ok","version":"v1"}
```

## 5. Access the UI

```bash
kubectl port-forward service/go-lab-ui 8080:80 --address=0.0.0.0
```

Open `http://<VM_IP>:8080` in your browser.

## Teardown

```bash
helm uninstall go-lab-ui go-lab-api go-lab-db
cd infra/k8s-node1 && terraform destroy -auto-approve
```
