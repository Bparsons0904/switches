# Base image for building the Go application
FROM golang:latest AS base
WORKDIR /app

# Cache dependencies by copying go.mod and go.sum separately
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire application source code
COPY . .

# Stage for Templ generation
FROM ghcr.io/a-h/templ:latest AS generate-stage
WORKDIR /app
COPY --from=base /app /app
USER root
RUN ["templ", "generate"]

# Build the Go application and migrator
FROM golang:latest AS build-stage
WORKDIR /app
COPY --from=generate-stage /app /app
RUN CGO_ENABLED=0 GOOS=linux go build -buildvcs=false -o /app/main ./cmd/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -buildvcs=false -o /app/migrator ./migrator/migrator.go

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
