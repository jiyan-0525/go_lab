package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var clickCount = 0
var serverPort = "8010"

func main() {
	// Read PORT from environment variable (default: 8010)
	if envPort := os.Getenv("PORT"); envPort != "" {
		serverPort = envPort
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlePage)
	mux.HandleFunc("/clicked", handleClicked)

	server := http.Server{
		Addr:         ":" + serverPort,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Printf("🚀 Server started on http://localhost:%s (PORT=%s)\n", serverPort, serverPort)
	log.Fatal(server.ListenAndServe())
}

func handlePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(200)

	page := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Go + HTMX Modern Dashboard</title>
    <!-- HTMX CDN -->
    <script src="https://unpkg.com/htmx.org@1.9.12"></script>
    <style>
        :root {
            --go-blue: #00ADD8;
            --go-dark: #007d9c;
            --bg-color: #f1f5f9;
            --card-bg: #ffffff;
            --text-main: #0f172a;
            --text-muted: #64748b;
            --border-color: #e2e8f0;
            --success: #10b981;
        }

        * {
            box-sizing: border-box;
            margin: 0;
            padding: 0;
        }

        body {
            font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif;
            background: linear-gradient(135deg, #f8fafc 0%, #e2e8f0 100%);
            min-height: 100vh;
            display: flex;
            align-items: center;
            justify-content: center;
            padding: 20px;
            color: var(--text-main);
        }

        .container {
            width: 100%;
            max-width: 480px;
            background: var(--card-bg);
            border-radius: 16px;
            box-shadow: 0 20px 25px -5px rgba(0, 0, 0, 0.08), 0 8px 10px -6px rgba(0, 0, 0, 0.04);
            border: 1px solid var(--border-color);
            overflow: hidden;
            transition: transform 0.2s ease;
        }

        .header {
            padding: 28px 28px 20px 28px;
            border-bottom: 1px solid var(--border-color);
            display: flex;
            align-items: center;
            justify-content: space-between;
        }

        .brand {
            display: flex;
            align-items: center;
            gap: 12px;
        }

        .logo-icon {
            width: 42px;
            height: 42px;
            background: linear-gradient(135deg, var(--go-blue) 0%, var(--go-dark) 100%);
            border-radius: 10px;
            display: flex;
            align-items: center;
            justify-content: center;
            font-size: 20px;
            color: white;
            box-shadow: 0 4px 6px -1px rgba(0, 173, 216, 0.3);
        }

        .title h1 {
            font-size: 1.15rem;
            font-weight: 700;
            color: var(--text-main);
        }

        .title p {
            font-size: 0.85rem;
            color: var(--text-muted);
        }

        .status-badge {
            display: inline-flex;
            align-items: center;
            gap: 6px;
            padding: 4px 10px;
            border-radius: 9999px;
            background: #ecfdf5;
            color: #065f46;
            font-size: 0.75rem;
            font-weight: 600;
        }

        .status-dot {
            width: 7px;
            height: 7px;
            background: var(--success);
            border-radius: 50%;
            box-shadow: 0 0 0 0 rgba(16, 185, 129, 0.7);
            animation: pulse 2s infinite;
        }

        @keyframes pulse {
            0% { transform: scale(0.95); box-shadow: 0 0 0 0 rgba(16, 185, 129, 0.7); }
            70% { transform: scale(1); box-shadow: 0 0 0 6px rgba(16, 185, 129, 0); }
            100% { transform: scale(0.95); box-shadow: 0 0 0 0 rgba(16, 185, 129, 0); }
        }

        .content {
            padding: 28px;
            display: flex;
            flex-direction: column;
            gap: 20px;
        }

        .btn-action {
            display: flex;
            align-items: center;
            justify-content: center;
            gap: 8px;
            width: 100%;
            padding: 14px 20px;
            background: linear-gradient(135deg, var(--go-blue) 0%, var(--go-dark) 100%);
            color: white;
            border: none;
            border-radius: 10px;
            font-size: 1rem;
            font-weight: 600;
            cursor: pointer;
            box-shadow: 0 4px 12px rgba(0, 173, 216, 0.25);
            transition: all 0.2s ease;
        }

        .btn-action:hover {
            transform: translateY(-1px);
            box-shadow: 0 6px 16px rgba(0, 173, 216, 0.35);
            filter: brightness(1.05);
        }

        .btn-action:active {
            transform: translateY(1px);
            box-shadow: 0 2px 6px rgba(0, 173, 216, 0.2);
        }

        #result-box {
            background: #f8fafc;
            border: 1px dashed var(--border-color);
            border-radius: 12px;
            padding: 20px;
            min-height: 110px;
            display: flex;
            align-items: center;
            justify-content: center;
            transition: all 0.3s ease;
        }

        .idle-state {
            color: var(--text-muted);
            font-size: 0.9rem;
            display: flex;
            align-items: center;
            gap: 8px;
        }

        .active-card {
            width: 100%;
            text-align: center;
            animation: fadeIn 0.3s ease;
        }

        @keyframes fadeIn {
            from { opacity: 0; transform: scale(0.97); }
            to { opacity: 1; transform: scale(1); }
        }

        .counter-badge {
            font-size: 2.2rem;
            font-weight: 800;
            color: var(--go-blue);
            line-height: 1.1;
        }

        .counter-label {
            font-size: 0.8rem;
            text-transform: uppercase;
            letter-spacing: 0.05em;
            color: var(--text-muted);
            margin-top: 4px;
        }

        .time-pill {
            display: inline-block;
            margin-top: 12px;
            font-size: 0.8rem;
            background: #e0f2fe;
            color: #0369a1;
            padding: 3px 10px;
            border-radius: 9999px;
            font-weight: 500;
        }

        .footer {
            padding: 16px 28px;
            background: #f8fafc;
            border-top: 1px solid var(--border-color);
            display: flex;
            justify-content: space-between;
            align-items: center;
            font-size: 0.78rem;
            color: var(--text-muted);
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <div class="brand">
                <div class="logo-icon">🐹</div>
                <div class="title">
                    <h1>Go + HTMX</h1>
                    <p>Reactive UI without React</p>
                </div>
            </div>
            <div class="status-badge">
                <span class="status-dot"></span>
                Online
            </div>
        </div>

        <div class="content">
            <!-- HTMX trigger: POST to /clicked, swap into #result-box -->
            <button class="btn-action" hx-post="/clicked" hx-target="#result-box">
                <span>⚡ Click Me to Interact</span>
            </button>

            <div id="result-box">
                <div class="idle-state">
                    <span>✨ Ready. Click the button above!</span>
                </div>
            </div>
        </div>

        <div class="footer">
            <span>Powered by <strong>net/http</strong> & <strong>HTMX</strong></span>
            <span>Port: ` + serverPort + `</span>
        </div>
    </div>
</body>
</html>`

	w.Write([]byte(page))
}

func handleClicked(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	clickCount++
	now := time.Now().Format("15:04:05")

	responseHTML := fmt.Sprintf(`
        <div class="active-card">
            <div class="counter-badge">%d</div>
            <div class="counter-label">Server Invocations</div>
            <div class="time-pill">🕒 Last Response: %s</div>
        </div>
    `, clickCount, now)

	w.Write([]byte(responseHTML))
}
