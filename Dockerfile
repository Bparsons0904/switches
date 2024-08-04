FROM golang:1.22 AS base

# Create a stage for downloading dependencies
FROM base AS deps
WORKDIR /app
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
  go mod download

# Stage for Templ generation
FROM ghcr.io/a-h/templ:latest AS generate-stage
WORKDIR /app
COPY --from=deps /app /app
COPY . .
USER root
RUN ["templ", "generate"]

# Build the Go application and migrator
FROM base AS build-stage
WORKDIR /app
COPY --from=generate-stage /app /app
RUN --mount=type=cache,target=/go/pkg/mod \
  --mount=type=cache,target=/root/.cache/go-build \
  CGO_ENABLED=0 GOOS=linux go build -buildvcs=false -o /app/main ./cmd/main.go && \
  CGO_ENABLED=0 GOOS=linux go build -buildvcs=false -o /app/migrator ./migrator/migrator.go

# Final deployment stage
FROM alpine:latest AS deploy-stage
WORKDIR /app
COPY --from=build-stage /app/main /app/main
COPY --from=build-stage /app/migrator /app/migrator
COPY --from=build-stage /app/static /app/static

# Install necessary packages
RUN apk update && apk add --no-cache bind-tools busybox-extras openssl

# Create a non-root user and switch to that user
RUN addgroup -S nonroot && adduser -S nonroot -G nonroot

# Set permissions for directories and switch to the non-root user
RUN mkdir -p /app/assets && \
  chown -R nonroot:nonroot /app && \
  chmod -R 755 /app
USER nonroot

# Expose the application port
EXPOSE 3030

# Set the entry point for the application
ENTRYPOINT ["/app/main"]
