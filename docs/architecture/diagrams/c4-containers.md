# C4 Level 2 — Containers

Shows the three runtime components inside the platform and how they communicate.

```mermaid
C4Container
    title Container Diagram — lab_web_platform

    Person(user, "User", "Browser")

    System_Boundary(k8s, "Kubernetes Cluster (k8s-lab-n1)") {
        Container(ui, "go-lab-ui", "nginx", "Serves static HTML/JS and proxies /api/* to go-lab-api")
        Container(api, "go-lab-api", "Go / Fiber v2", "REST API: GET /healthz, POST /api/ping — increments click counter")
        ContainerDb(db, "go-lab-db", "PostgreSQL 16", "Stores button_stats table with click counter")
    }

    Rel(user, ui, "HTTP GET /", "HTTP :80")
    Rel(user, ui, "POST /api/ping", "HTTP :80")
    Rel(ui, api, "Proxy /api/*", "HTTP :8080")
    Rel(ui, api, "Proxy /healthz", "HTTP :8080")
    Rel(api, db, "SELECT / UPDATE button_stats", "TCP :5432")
```
