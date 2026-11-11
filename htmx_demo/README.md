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
cd htmx_demo

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

We use a **Multi-Stage Dockerfile** to produce an ultra-lightweight (~19MB) production image:

1. **Build the Docker Image**:
   ```bash
   cd htmx_demo
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
├── go.mod         # Module definition (standard library only, no dependencies)
├── Dockerfile     # Multi-stage build producing a ~19MB static image
├── .dockerignore  # Keeps .git and docs out of the build context
└── README.md      # Project documentation
```

> **Routing (Go 1.22+)**: routes are registered with method-aware patterns — `GET /{$}` for the page (exact root only) and `POST /clicked` for the interaction. `ServeMux` answers wrong methods with `405 Method Not Allowed` on its own, so a plain visit to `/clicked` in the browser can no longer bump the counter.

---

## 💡 Why Go + HTMX?

For backend and systems developers, writing complex React/Vue frontend setups often introduces excessive tooling overhead (JavaScript ecosystem, state management, bundling, API synchronization).

**HTMX brings the simplicity back**:
- State resides on the server (Go).
- The server renders HTML, which browsers natively understand.
- You get 90% of modern SPA interactivity with 10% of the complexity.

---

## 🔍 Technical Comparison: React SPA vs. Go + HTMX

The table below summarizes the core design trade-offs this demo is meant to illustrate.

| Aspect | Typical React SPA | This Demo |
| --- | --- | --- |
| **Where state lives** | Frontend JavaScript (Redux, hooks, etc.) | Server side, in Go (`var clickCount`) |
| **What travels over the wire** | JSON payloads, rendered by the client | HTML fragments, rendered by the server |
| **Build toolchain** | npm / webpack / Vite | None — `go run main.go` just works |
| **Dependencies** | Large JS ecosystem | Only the standard library: `net/http`, `fmt`, `time` |
| **How the UI updates** | Client-side virtual DOM re-render | HTMX swaps a partial DOM fragment into `#result-box` |
| **Routing & data fetching** | Separate HTTP client + state synchronization | A plain `POST /clicked` returning ready-to-use markup |

### Key takeaways

- **Server-owned state**: the click counter lives in a single Go variable, so there is no client-side state store to keep in sync.
- **HTML as the API contract**: handlers return HTML snippets (`<div class="active-card">…</div>`) instead of JSON, and the browser already knows how to display them.
- **No build pipeline**: a single Go binary serves the page, the styles, and the interaction logic.
- **Progressive enhancement**: HTMX is loaded from a CDN, so the application itself is a single `.go` file — everything else in the directory is packaging (module file, Dockerfile, docs).

> ⚠️ **Demo, not production code**: `clickCount` is a plain global variable, which is not safe under concurrent requests or across multiple replicas, and there is no authentication or persistence.
