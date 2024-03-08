# Use an official Golang runtime as the base image
FROM golang:1.18-alpine

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app/.

# Install build-essential package
RUN apk add --no-cache build-base


# Compile the Go binary with CGO enabled
RUN go build -o out

# Set the entrypoint to run the Go binary
ENTRYPOINT ["./out"]