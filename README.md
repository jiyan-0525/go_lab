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
├── go_exercises/           # 📚 Core Language Exercises & Algorithms
│   ├── go_by_Example/      # Go by Example fundamentals
│   │   ├── hello.go        # Hello, world entry point
│   │   ├── variables.go    # Variables, types & string indexing
│   │   ├── control.go      # if/else control flow
│   │   ├── for_loops.go    # Iterations, range over slices & maps
│   │   ├── switch.go       # switch statements & custom types
│   │   ├── funtions.go     # Functions, named returns & multiple returns
│   │   ├── slices.go       # Arrays, slices, append & sub-slicing
│   │   ├── maps.go         # Key-value dictionary operations
│   │   ├── pointers.go     # Memory addresses & pointer semantics
│   │   └── rune.go         # Runes & Unicode string handling
│   └── go_by_web/          # Web, cache, ORM exercises (planned)
│
├── go_server/              # 🌐 Backend Microservice & Docker Practices
│   ├── main.go             # Standard library HTTP server
│   ├── Dockerfile          # Self-compiling multi-stage Dockerfile
│   ├── go.mod              # Go module definition
│   ├── README.md           # Server & containerization notes
│   └── envcreat/           # 12-Factor dynamic environment config sub-lab
│
└── htmx_demo/              # 🐹 Reactive Web Application
    ├── main.go             # Go server with dynamic PORT & inline HTMX template
    ├── Dockerfile          # Multi-stage production build (~12MB image)
    ├── go.mod              # Go module definition
    └── README.md           # HTMX interactive guide
```

---

## 🔄 Automated CI/CD Pipeline (Azure DevOps)

The repository uses **Azure DevOps Pipelines** ([`azure-pipelines.yml`](./azure-pipelines.yml)) to achieve fully automated continuous integration and continuous delivery (CI/CD) to **Docker Hub**.

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
| [**`go_exercises/`**](./go_exercises) | Core language syntax, data structures, and algorithms | Variables, Control Flow, Functions, Slices, Maps, Pointers, Runes | *N/A (Library Code)* |
| [**`go_server/`**](./go_server) | Microservice HTTP server & cloud-native dynamic configuration | Go, Docker, 12-Factor App (`PORT`) | `jiyanwang/goserver` |
| [**`htmx_demo/`**](./htmx_demo) | Modern single-page interactive UI without Node.js or frontend bundlers | Go `net/http`, HTMX, Responsive CSS | `jiyanwang/htmx-dashboard` |

---

## 🚀 Quick Start & Docker Hub Deployment

Both microservices are automatically built, tested, and published to **Docker Hub** via the Azure DevOps CI/CD pipeline.

### Option 1: Direct Cloud Deployment from Docker Hub

Run the pre-built, production-ready images anywhere with Docker:

#### 🐹 1. Run HTMX Reactive Dashboard
```bash
# Pull the latest image
docker pull jiyanwang/htmx-dashboard:latest

# Run the container (accessible on http://localhost:8010)
docker run -d -p 8010:8010 --name htmx-dashboard jiyanwang/htmx-dashboard:latest

# Or run with dynamic custom port:
docker run -d -e PORT=8999 -p 8999:8999 --name htmx-dashboard jiyanwang/htmx-dashboard:latest
```

#### 🌐 2. Run Go HTTP Server
```bash
# Pull the latest image
docker pull jiyanwang/goserver:latest

# Run the container (accessible on http://localhost:8020)
docker run -d -p 8020:8010 --name goserver jiyanwang/goserver:latest
```

---

### Option 2: Local Development

```bash
# 1. Run HTMX Dashboard locally:
cd htmx_demo
go run main.go

# 2. Run Go Server locally:
cd ../go_server
go run main.go
```
