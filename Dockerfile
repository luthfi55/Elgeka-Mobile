# Use an official Golang runtime as the base image
FROM golang:1.20-alpine

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app/.

# Install build-essential package
RUN apk add --no-cache build-base

# Print the contents of the go.mod file
RUN cat go.mod

# Set the CGO_ENABLED environment variable to 1
ENV CGO_ENABLED 1

# Compile the Go binary with CGO enabled
RUN go build -o out 

# Set the entrypoint to run the Go binary
ENTRYPOINT ["./out"]