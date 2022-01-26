#FROM golang:1.12.0-alpine3.9
FROM golang:1.17-buster AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app

ARG GOOS=linux
ARG GOARCH=amd64
ARG GOFLAGS=-mod=vendor
ARG CGO_ENABLED=0
ARG GOOS=linux
## Add this go mod download command to pull in any dependencies
## Our project will now successfully build with the necessary go libraries included.
RUN  go build  -o main .

FROM debian:buster-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/main /app/main

# Run the web service on container startup.
CMD ["/app/main"]




# https://hub.docker.com/_/golang
#FROM golang:1.17-buster as builder
#
## Create and change to the app directory.
#WORKDIR /app
#
## Retrieve application dependencies.
## This allows the container build to reuse cached dependencies.
## Expecting to copy go.mod and if present go.sum.
#COPY go.* ./
#RUN go mod download
#
## Copy local code to the container image.
#COPY . ./
#
## Build the binary.
#RUN go build -v -o server
#
## Use the official Debian slim image for a lean production container.
## https://hub.docker.com/_/debian
## https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
#FROM debian:buster-slim
#RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
#    ca-certificates && \
#    rm -rf /var/lib/apt/lists/*
#
## Copy the binary to the production image from the builder stage.
#COPY --from=builder /app/server /app/server
#
## Run the web service on container startup.
#CMD ["/app/server"]
#
## [END run_helloworld_dockerfile]
## [END cloudrun_helloworld_dockerfile]
