# ---------- Stage 1: Build the Go binary ----------
FROM golang:1.23.2-alpine AS builder

RUN apk add --no-cache git curl

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Download golang-migrate CLI
ENV MIGRATE_VERSION=v4.18.3
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/${MIGRATE_VERSION}/migrate.linux-amd64.tar.gz \
    | tar xvz \
    && chmod +x migrate \
    && mv migrate /usr/local/bin/migrate

# IMPORTANT: Build for Linux and disable CGO
RUN go build -ldflags '-w -s' -a -o main .

CMD ["/app/main"]

EXPOSE 4000

