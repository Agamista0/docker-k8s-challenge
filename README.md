🚀 DevOps Challenge -- EYouth X DEPI Tech
========================================

This repository contains my solution to the **Docker & Kubernetes Challenge**.

📌 Challenge Overview
---------------------

The goal was to:

1.  Containerize a Go application with Redis.

2.  Optimize Docker image size.

3.  Deploy the application on Kubernetes.

4.  Troubleshoot and fix issues encountered during deployment.

* * * * *

🐳 Docker Setup
---------------

-   Built **two containers**:

    1.  Go Application

    2.  Redis

-   Connected them using Docker's internal network.

✅ Proof of working setup included in screenshots.

### Image Optimization

-   **Before:** `1.03 GB`

-   **After:** `24.5 MB`

-   Achieved by using a **multi-stage Docker build**.

* * * * *

☸️ Kubernetes Deployment
------------------------

Deployment was done in **3 steps**:

1.  **Namespaces** -- Created and applied namespaces for better organization.

2.  **Redis** -- Configured storage, deployment, and verified functionality.

3.  **Go App** -- Deployed and verified application functionality.

* * * * *

⚡ Issue Encountered
-------------------

-   **Problem:** Readiness & liveness probes were hitting `/`, which increased the visit counter incorrectly.

-   **Solution:** Added a dedicated `/health` endpoint for probes.

-   **Outcome:** Health checks no longer interfere with real visit counts.

* * * * *

📂 Repository Structure
-----------------------

`docker-k8s-challenge/
├── Dockerfile
├── k8s/
│   ├── namespace.yaml
│   ├── redis-deployment.yaml
│   ├── redis-service.yaml
│   ├── app-deployment.yaml
│   ├── app-service.yaml
└── README.md`
