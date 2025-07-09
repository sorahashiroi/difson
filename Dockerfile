FROM golang:1-bullseye AS builder
WORKDIR /work

ARG CGO_ENABLED=0
COPY . .
RUN go build -o difson main.go

FROM alpine:latest
ARG VERSION=0.5.1

LABEL org.opencontainers.image.source=https://github.com/sorahashiroi/difson \
      org.opencontainers.image.version=${VERSION} \
      org.opencontainers.image.title=difson \
      org.opencontainers.image.description="A tool for diffing JSON files"

RUN addgroup -S nonroot && \
    adduser -S -D -H -h /workdir nonroot nonroot && \
    mkdir -p /workdir

COPY --from=builder /work/difson /opt/difson/difson
RUN apk --no-cache add ca-certificates

WORKDIR /workdir
USER nonroot

ENTRYPOINT [ "/opt/difson/difson" ]