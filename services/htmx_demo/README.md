# 🐹 Go + HTMX Reactive Dashboard

A lightweight, modern, and reactive web application built with **pure Go standard library (`net/http`)** and **[HTMX](https://htmx.org/)**.

> **No Node.js. No npm. No Webpack/Vite. No React/Vue.**  
> Enjoy the fluid responsiveness of a Single Page Application (SPA) driven directly by a single Go binary.

---

## 📸 Overview & Features

- **⚡ Zero Build Pipeline**: Runs straight out of the box with `go run main.go`.
- **🔄 Seamless Partial DOM Swapping**: Powered by HTMX, the client triggers async HTTP POST requests and dynamically updates page components without refreshing the browser.
- **🎨 Modern Aesthetic UI**:
  - Gopher cyan gradient theme (`#00ADD8`)
  - Elevated card layout with soft shadows and border radiuses
  - Pulsing live server status badge (`Online`)
  - Smooth micro-animations on interactive triggers
- **🛡️ 100% Standard Library**: Uses only Go's built-in `net/http`, `fmt`, and `time` packages.

---

## 🚀 Quick Start

### Option A: Local Run
```bash
cd services/htmx_demo

# Run with default port 8010:
go run main.go

# Or run with a custom port:
PORT=9000 go run main.go
```
Visit `http://localhost:8010` (or your custom port) in your browser.

---

### Option B: Pull & Run directly from Docker Hub (Automated via Azure DevOps)

This image is automatically built and published by the root Azure DevOps CI/CD pipeline:

```bash
# Pull the latest image from Docker Hub
docker pull jiyanwang/htmx-dashboard:latest

# Run the container
docker run -d -p 8010:8010 --name htmx-app jiyanwang/htmx-dashboard:latest

# Or run with dynamic custom port (e.g. 9000):
docker run -d -e PORT=9000 -p 9000:9000 --name htmx-app jiyanwang/htmx-dashboard:latest
```

---

### Option C: Build and Run Docker Locally

We use a **Multi-Stage Dockerfile** to produce an ultra-lightweight (~12MB) production image:

1. **Build the Docker Image**:
   ```bash
   cd services/htmx_demo
   docker build -t htmx-dashboard:latest .
   ```

2. **Run with Default Port (8010)**:
   ```bash
   docker run -d -p 8010:8010 --name htmx-app htmx-dashboard:latest
   ```

3. **Run with Custom Dynamic Port (e.g., 9090)**:
   ```bash
   docker run -d -e PORT=9090 -p 9090:9090 --name htmx-custom htmx-dashboard:latest
   ```

---

## 🛠️ How It Works

```text
┌───────────────────────┐                    ┌─────────────────────────┐
│     Client Browser    │                    │     Go HTTP Server      │
│                       │                    │    (Port: 8010)         │
│  <button              │  POST /clicked     │                         │
│    hx-post="/clicked" ├───────────────────>│  handleClicked(w, r)    │
│    hx-target="#box">  │                    │  • Increments clickCount│
│                       │  Partial HTML      │  • Formats timestamp    │
│  DOM swapped smoothly │<───────────────────┤  • Returns HTML snippet │
│  (No full-page reload)│                    │                         │
└───────────────────────┘                    └─────────────────────────┘
```

1. **Initial Page Load (`/`)**:
   - The server delivers the HTML shell containing embedded CSS and the HTMX CDN script tag.
2. **Interactive Action (`hx-post="/clicked"`)**:
   - When the user clicks the button, HTMX intercepts the event and dispatches an asynchronous HTTP POST request to `/clicked`.
3. **Partial HTML Swap (`hx-target="#result-box"`)**:
   - The Go handler computes the new counter value and timestamp, formats a clean HTML fragment (`<div class="active-card">...</div>`), and returns it.
   - HTMX swaps this fragment directly into `#result-box`.

---

## 📂 Project Structure

```text
htmx_demo/
├── main.go        # HTTP server, routes, inline HTML shell & handlers
└── README.md      # Project documentation
```

---

## 💡 Why Go + HTMX?

For backend and systems developers, writing complex React/Vue frontend setups often introduces excessive tooling overhead (JavaScript ecosystem, state management, bundling, API synchronization).

**HTMX brings the simplicity back**:
- State resides on the server (Go).
- The server renders HTML, which browsers natively understand.
- You get 90% of modern SPA interactivity with 10% of the complexity.
