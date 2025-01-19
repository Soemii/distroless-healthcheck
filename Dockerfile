ARG DISTROLESS_IMAGE=gcr.io/distroless/static-debian12
FROM golang:latest AS build
WORKDIR /app
COPY . .
RUN go mod download && make build

FROM ${DISTROLESS_IMAGE}
WORKDIR /
COPY --from=build /app/tmp/healthcheck /usr/local/bin/healthcheck