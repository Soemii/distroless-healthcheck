ARG DISTROLESS_IMAGE=static-debian12
FROM golang:latest AS build
WORKDIR /app
COPY . .
RUN go mod download && make build

FROM gcr.io/distroless/${DISTROLESS_IMAGE}
WORKDIR /
COPY --from=build /app/tmp/healthcheck /usr/local/bin/healthcheck