# Use an official Golang runtime as a parent image
FROM golang:1.22-alpine AS builder

RUN apk add --no-cache gcc musl-dev

# Set the working directory inside the container
WORKDIR /app

# Copy the go mod and sum files
COPY bin/go.mod ./
COPY bin/go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the entire project and build it
# This assumes that your Dockerfile is in the root of your Go project
# and that all your Go files are within this directory or subdirectories
COPY ./bin .

# Build the Go app
RUN CGO_ENABLED=1 GOOS=linux go build -v -o goserver

# Use the alpine base image for the runtime container
FROM alpine:latest  
WORKDIR /root/

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/goserver .

# Command to run the executable
CMD ["./goserver"]

VOLUME ["/var/www/db", "/var/www/web/templates", "/var/www/static"]