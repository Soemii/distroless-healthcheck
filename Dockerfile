FROM golang:latest AS build
WORKDIR /app
COPY . .
RUN go mod download && make build

ARG DISTROLESS_IMAGE=gcr.io/distroless/static-debian12

FROM $DISTROLESS_IMAGE
LABEL org.opencontainers.image.source=https://github.com/Soemii/distroless-healthcheck
WORKDIR /
COPY --from=build /app/tmp/healthcheck /usr/local/bin/healthcheck