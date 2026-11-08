# 🧪 Go Lab

[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![Docker](https://img.shields.io/badge/Docker-Containerized-2496ED?style=for-the-badge&logo=docker&logoColor=white)](https://www.docker.com/)
[![Azure DevOps](https://img.shields.io/badge/Azure_DevOps-CI%2FCD_Pipelines-0078D7?style=for-the-badge&logo=azuredevops&logoColor=white)](https://azure.microsoft.com/en-us/products/devops/)
[![HTMX](https://img.shields.io/badge/HTMX-Reactive_UI-3366CC?style=for-the-badge&logo=htmx&logoColor=white)](https://htmx.org/)

A hands-on Go engineering and learning laboratory featuring language fundamentals, micro-services containerization, reactive fullstack web applications, and automated **Azure DevOps CI/CD** cloud delivery.

---

## 🏗️ Repository Architecture

This repository is organized as a structured **Monorepo**:

```text
go_lab/
├── azure-pipelines.yml     # 🚀 Root Azure DevOps CI/CD pipeline (Multi-job)
├── .gitignore              # 🛡️ Global binary and build artifact filters
├── README.md               # 📖 Repository documentation
│
├── htmx_demo/              # 🐹 Reactive Web Application
│   ├── main.go             # Go server with dynamic PORT & inline HTMX template
│   ├── Dockerfile          # Multi-stage production build (~12MB image)
│   ├── go.mod              # Go module definition
│   └── README.md           # HTMX interactive guide
│
├── go_server/              # 🌐 Backend Microservice & Docker Practices
│   ├── main.go             # Standard library HTTP server
│   ├── Dockerfile          # Self-compiling multi-stage Dockerfile
│   ├── go.mod              # Go module definition
│   ├── README.md           # Server & containerization notes
│   └── envcreat/           # 12-Factor dynamic environment config sub-lab
│
└── go_exercises/           # 📚 Core Language Exercises & Algorithms
    ├── for_loops.go        # Iterations, control flow & range slices
    ├── slices.go           # Slice manipulation & dynamic arrays
    ├── maps.go             # Key-value dictionary operations
    ├── pointers.go         # Memory addresses & pointer semantics
    └── ...
```

---

## 🔄 Automated CI/CD Pipeline (Azure DevOps)

The repository uses **Azure DevOps Pipelines** ([`azure-pipelines.yml`](file:///home/jiyan/go/go_lab/azure-pipelines.yml)) to achieve fully automated continuous integration and continuous delivery (CI/CD) to **Docker Hub**.

### Pipeline Workflow

```text
  Git Push to main
        │
        ▼
┌────────────────────────────────────────────────────────┐
│             Azure DevOps CI/CD Pipeline                │
│                                                        │
│  Intelligent Path Filtering:                           │
│  - Triggers only when `go_server/**` or `htmx_demo/**` │
│    files change.                                       │
└───────────┬────────────────────────────────┬───────────┘
            │                                │
            ▼                                ▼
┌─────────────────────────┐      ┌─────────────────────────┐
│ Task 1: Build & Push    │      │ Task 2: Build & Push    │
│ Target: go_server       │      │ Target: htmx_demo       │
│ Multi-stage Dockerfile  │      │ Multi-stage Alpine Build│
└───────────┬─────────────┘      └───────────┬─────────────┘
            │                                │
            ▼                                ▼
┌─────────────────────────┐      ┌─────────────────────────┐
│       Docker Hub        │      │       Docker Hub        │
│    jiyanwang/goserver   │      │ jiyanwang/htmx-dashboard│
│   (latest, $(BuildId))  │      │   (latest, $(BuildId))  │
└─────────────────────────┘      └─────────────────────────┘
```

### Key DevOps Highlights:
- **Intelligent Triggering**: Path-based triggers avoid redundant builds when updating exercises or documentation.
- **Hermetic Multi-Stage Builds**: Compilation happens entirely inside the Docker build container (`golang:1.22-alpine`), ensuring consistent, reproducible builds across any runner environment without pre-compiled binaries in Git.
- **Minimal Image Footprint**: Final runtime containers run on minimal scratch/Alpine/Debian bases, keeping production images down to ~12MB.

---

## 📁 Modules Breakdown

| Module | Purpose | Key Technologies | Docker Hub Image |
|---|---|---|---|
| [**`htmx_demo/`**](file:///home/jiyan/go/go_lab/htmx_demo) | Modern single-page interactive UI without Node.js or frontend bundlers | Go `net/http`, HTMX, Responsive CSS | `jiyanwang/htmx-dashboard` |
| [**`go_server/`**](file:///home/jiyan/go/go_lab/go_server) | Microservice HTTP server & cloud-native dynamic configuration | Go, Docker, 12-Factor App (`PORT`) | `jiyanwang/goserver` |
| [**`go_exercises/`**](file:///home/jiyan/go/go_lab/go_exercises) | Core language syntax, data structures, and algorithms | Control Flow, Slices, Maps, Pointers | *N/A (Library Code)* |

---

## 🚀 Quick Start

### 1. Run the Reactive HTMX Dashboard
```bash
cd htmx_demo

# Run locally:
go run main.go

# Or run with custom dynamic port:
PORT=9000 go run main.go
```
Visit `http://localhost:8010` (or `http://localhost:9000`).

### 2. Pull & Run from Docker Hub
Once Azure DevOps finishes pushing the images:
```bash
# Run the HTMX Dashboard container
docker run -d -p 8010:8010 --name htmx-dashboard jiyanwang/htmx-dashboard:latest

# Run the Go Server container
docker run -d -p 8020:8010 --name goserver jiyanwang/goserver:latest
```
