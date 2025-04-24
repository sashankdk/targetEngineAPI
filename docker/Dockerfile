# Stage 1: Build
FROM golang:1.21 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o targetApi ./cmd/main.go

# Stage 2: Run
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/targetApi .
CMD ["./targetApi"]
