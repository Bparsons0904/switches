FROM golang:latest AS fetch-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

FROM ghcr.io/a-h/templ:latest AS generate-stage
WORKDIR /app
COPY --from=fetch-stage /app /app
COPY . /app
USER root
RUN ["templ", "generate"]

FROM golang:latest AS build-stage
WORKDIR /app
COPY --from=generate-stage /app /app
RUN CGO_ENABLED=0 GOOS=linux go build -buildvcs=false -o /app/main ./cmd/main.go

FROM alpine:latest AS deploy-stage
WORKDIR /app
COPY --from=build-stage /app/main /app/main
COPY --from=build-stage /app/static /app/static
RUN apk update && apk add --no-cache bind-tools busybox-extras openssl
RUN addgroup -S nonroot && adduser -S nonroot -G nonroot

USER nonroot
EXPOSE 3076
ENTRYPOINT ["/app/main"]
