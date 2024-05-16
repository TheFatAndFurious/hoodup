# Use an official Golang runtime as a parent image
FROM golang:1.22-alpine AS builder

# Install necessary dependencies
RUN apk add --no-cache gcc musl-dev

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=1 GOOS=linux go build -v -o goserver ./cmd/server/main.go

# Use the alpine base image for the runtime container
FROM alpine:latest
WORKDIR /app

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/goserver .

# Command to run the executable
CMD ["./goserver"]

# Define volumes for persistence
VOLUME ["/app/static", "/app/db", "/app/public", "/app/web/templates"]

