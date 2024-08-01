# Base image for building the Go application
FROM golang:latest AS base
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Stage for Templ generation
FROM ghcr.io/a-h/templ:latest AS generate-stage
WORKDIR /app
COPY --from=base /app /app
USER root
RUN ["templ", "generate"]

# Build the Go application
FROM golang:latest AS build-stage
WORKDIR /app
COPY --from=generate-stage /app /app
RUN CGO_ENABLED=0 GOOS=linux go build -buildvcs=false -o /app/main ./cmd/main.go

# Final deployment stage
FROM alpine:latest AS deploy-stage
WORKDIR /app
COPY --from=build-stage /app/main /app/main
COPY --from=build-stage /app/static /app/static

# Install necessary packages
RUN apk update && apk add --no-cache bind-tools busybox-extras openssl

# Create necessary directories with correct permissions
RUN mkdir -p /app/assets && \
    chown -R nonroot:nonroot /app/assets && \
    chmod -R 777 /app/assets

# Create a non-root user and switch to that user
RUN addgroup -S nonroot && adduser -S nonroot -G nonroot
USER nonroot

# Expose the application port
EXPOSE 3030

# Set the entry point for the application
ENTRYPOINT ["/app/main"]
