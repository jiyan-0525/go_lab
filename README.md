# 🧪 Go Lab

A hands-on Go development and learning laboratory featuring core language exercises, Dockerized backend services, and modern reactive web interfaces.

---

## 📁 Repository Overview

| Module | Description | Key Tech |
|---|---|---|
| [**`htmx_demo/`**](file:///home/jiyan/go/go_lab/htmx_demo) | Reactive dashboard without frontend build tools | Go `net/http`, HTMX, Modern CSS |
| [**`go_server/`**](file:///home/jiyan/go/go_lab/go_server) | Dockerized HTTP server & 12-Factor environment configuration | Go, Docker, Cloud-Native Env |
| [**`go_exercises/`**](file:///home/jiyan/go/go_lab/go_exercises) | Core language syntax, data structures, and algorithms | Control Flow, Slices, Maps, Pointers |

---

## 🚀 Quick Navigation

- **Try Go + HTMX Dashboard**:
  ```bash
  cd htmx_demo
  go run main.go
  # Open http://localhost:8010
  ```

- **Run Dockerized Go Server**:
  ```bash
  cd go_server
  docker build . -t goserver:latest
  docker run -p 8010:8010 goserver
  ```
