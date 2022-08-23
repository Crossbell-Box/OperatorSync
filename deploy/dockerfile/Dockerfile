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
RUN go build -o ./build/server ./app/server

FROM alpine:latest AS RUNNER

WORKDIR /app

COPY --from=BUILDER /app/build/* /usr/local/bin/
