# Deployment Diagram

Shows the full deployment pipeline from code to running cluster.

```mermaid
flowchart TD
    subgraph Dev["Developer Machine"]
        code["Source Code\n(lab_web_platform)"]
        tf["Terraform\ninfra/k8s-node1"]
    end

    subgraph GitHub
        repo["GitHub Repository\nmain branch"]
        actions["GitHub Actions\nrelease.yaml"]
        releases["GitHub Releases\nui-v*.tar.gz\ngo-lab-api-*.tgz\ngo-lab-db-*.tgz"]
        ghcr["GHCR\nghcr.io/dplaksa/go-lab-api"]
    end

    subgraph Multipass["Multipass VM (k8s-lab-n1 — 10.97.0.144)"]
        subgraph K8s["Kubernetes Cluster"]
            ui_pod["go-lab-ui Pod\nnginx + initContainer"]
            api_pod["go-lab-api Pod\nGo binary"]
            db_pod["go-lab-db StatefulSet\nPostgreSQL 16"]
            pvc["PersistentVolumeClaim\nlocal-path 1Gi"]
        end
    end

    code -->|git push| repo
    repo -->|tag ui-v* / api-v* / db-v*| actions
    actions -->|tar.gz + helm .tgz| releases
    actions -->|docker push| ghcr

    tf -->|multipass launch + cloud-init| Multipass
    releases -->|initContainer wget| ui_pod
    ghcr -->|imagePull| api_pod
    db_pod --> pvc

    ui_pod -->|proxy /api/*| api_pod
    api_pod -->|SQL| db_pod
```

## Infrastructure

| Component | Technology | Provisioner |
|---|---|---|
| VM | Ubuntu 26.04, 3 vCPU, 6GB RAM, 30GB disk | Terraform + Multipass |
| Kubernetes | v1.31, kubeadm | cloud-init |
| CNI | Flannel (VXLAN, 10.244.0.0/16) | cloud-init |
| Storage | local-path-provisioner | cloud-init |
| Package manager | Helm v3 | manual (planned: Flux CD) |
