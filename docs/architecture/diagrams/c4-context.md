# C4 Level 1 — System Context

Shows how the platform fits into its environment and who interacts with it.

```mermaid
C4Context
    title System Context — lab_web_platform

    Person(user, "User", "Opens the web page and clicks the button")

    System(platform, "lab_web_platform", "Single-page web application with a click counter backed by PostgreSQL")

    System_Ext(ghcr, "GitHub Container Registry", "Hosts the go-lab-api Docker image")
    System_Ext(github, "GitHub", "Source code, Helm chart releases and UI static bundle releases")

    Rel(user, platform, "Opens in browser", "HTTP")
    Rel(platform, ghcr, "Pulls container image", "HTTPS")
    Rel(platform, github, "Pulls UI release archive", "HTTPS")
```
