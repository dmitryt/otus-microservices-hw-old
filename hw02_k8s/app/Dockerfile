FROM golang:1.16 AS builder

# Set necessary environment variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux

WORKDIR /build
RUN go install github.com/go-task/task/v3/cmd/task@latest

COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN task build

# Build a small image
FROM alpine:latest

RUN apk add --no-cache bash

ARG RUN_MODE dev
ENV RUN_MODE $RUN_MODE

WORKDIR /app
COPY --from=builder /build/tmp/hw02-k8s-service ./
COPY --from=builder /build/conf ./conf
EXPOSE 8000

# Command to run
ENTRYPOINT ["./hw02-k8s-service"]