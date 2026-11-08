# ==========================================
# Stage 1: Build the Go server
# ==========================================
FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY main.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/goserver .

# ==========================================
# Stage 2: Minimal Runtime Image
# ==========================================
FROM debian:stable-slim

WORKDIR /app

COPY --from=builder /app/goserver /bin/goserver

EXPOSE 8010

CMD ["/bin/goserver"]
