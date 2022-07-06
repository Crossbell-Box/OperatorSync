FROM golang:alpine AS BUILDER

# Set the Current Working Directory inside the container
WORKDIR /app

# Install basic packages
RUN apk add \
    git gcc g++

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go mod download

# Build image
RUN go build -o ./build/worker ./app/worker

FROM alpine:latest AS RUNNER

WORKDIR /app

COPY --from=BUILDER /app/build/worker /app/worker

RUN chmod +x /app/worker
RUN ln -s /app/worker /usr/local/bin/worker

# This container exposes port 8080 to the outside world
EXPOSE 8080/tcp

# Environment variables
## RSSHub endpoints setting
ENV RSSHUB_STATEFUL=https://rsshub.app
ENV RSSHUB_STATELESS=https://rsshub.app
## Direct access http proxy URL, default disabled
#ENV PROXY_URL=http://localhost:4000
## IPFS Upload endpoint
ENV IPFS_ENDPOINT=https://ipfs-relay.crossbell.io/upload
## Message Queue connection
ENV MQ_CONNECTION_STRING=nats://localhost:4222
## Concurrency control
ENV CONCURRENCY_CONTROL_STATEFUL=10
ENV CONCURRENCY_CONTROL_STATELESS=50
ENV CONCURRENCY_CONTROL_DIRECT=100
## Work mode
ENV MODE=prod

# Run the executable
CMD ["worker"]
