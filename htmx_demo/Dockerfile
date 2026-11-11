# ==========================================
# Stage 1: Build the Go application
# ==========================================
# Pinned for reproducible builds (1.22.12 is the final patch of the 1.22 line).
FROM golang:1.22.12-alpine3.21 AS builder

WORKDIR /app

# This module has no third-party dependencies, so there is no `go mod download`
# step to cache: copy the source and build a clean, static binary directly.
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/server .

# ==========================================
# Stage 2: Minimal Production Image (~19MB)
# ==========================================
# Pinned for reproducible builds; bump deliberately when you want a newer base.
FROM alpine:3.24.4

WORKDIR /app

# Set default port (overridable at container runtime via -e PORT=...)
ENV PORT=8010

# Copy binary from the builder stage
COPY --from=builder /app/server /app/server

# Expose default port
EXPOSE 8010

# Start server
CMD ["/app/server"]
