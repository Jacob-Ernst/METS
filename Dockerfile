# Build the Go Binary.
FROM golang:1.18 as build_server
ENV CGO_ENABLED 0

# Copy the source code into the container.
COPY . /server

# Build the service binary.
WORKDIR /server/cmd/api
RUN go build

# Run the Go Binary in Alpine.
FROM alpine:3.16
COPY --from=build_server /server/cmd/api/api /service/server
WORKDIR /service

RUN apk add --update --no-cache \
    bash

LABEL org.opencontainers.image.title="server" \
      org.opencontainers.image.authors="Jacob Ernst <akumadude.je@gmail.com>" \
      org.opencontainers.image.source="https://gitlab.com/jacob-ernst/mets/-/tree/main/cmd/api"
