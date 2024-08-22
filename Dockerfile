FROM golang:1.22.5-alpine as builder

RUN mkdir /app
WORKDIR /app

RUN apk add --no-cache gcc g++ libc-dev librdkafka-dev pkgconf musl-dev make

# Copy files from current directory to docker image
COPY . .

# Download project dependencies
RUN go mod download

# Build Notification executable application
RUN go build -tags musl -o main cmd/main.go

# Base docker image for Golang applications
FROM golang:1.22.5-alpine

RUN apk add --no-cache make bash

RUN mkdir /app
WORKDIR /app

COPY --from=builder /app/.env .

# # Copy application executable files
COPY --from=builder /app/main /app/main

EXPOSE 3000

CMD ["./main"]
