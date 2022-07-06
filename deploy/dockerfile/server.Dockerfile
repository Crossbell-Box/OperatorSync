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
RUN go build -o ./build/server ./app/server

FROM alpine:latest AS RUNNER

WORKDIR /app

COPY --from=BUILDER /app/build/server /app/server

RUN chmod +x /app/server
RUN ln -s /app/server /usr/local/bin/server

# This container exposes port 8080 to the outside world
EXPOSE 8080/tcp

# Environment variables
ENV DATABASE_CONNECTION_STRING=postgres://postgres:postgres@localhost:5432/postgres
ENV REDIS_CONNECTION_STRING=redis://localhost:6379/0
ENV MQ_CONNECTION_STRING=nats://localhost:4222
ENV MODE=prod

# Run the executable
CMD ["server"]
