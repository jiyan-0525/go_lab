# ==========================================
# Stage 1: Build the Go application
# ==========================================
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Cache dependencies
COPY go.mod ./
RUN go mod download || true

# Copy source code and build a clean, static binary
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/server .

# ==========================================
# Stage 2: Minimal Production Image (~12MB)
# ==========================================
FROM alpine:latest

WORKDIR /app

# Set default port (overridable at container runtime via -e PORT=...)
ENV PORT=8010

# Copy binary from the builder stage
COPY --from=builder /app/server /app/server

# Expose default port
EXPOSE 8010

# Start server
CMD ["/app/server"]
