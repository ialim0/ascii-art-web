# Use the official Golang base image
FROM golang:1.19

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files to the container
COPY go.mod  ./

# Download and install the project dependencies
RUN go mod download

# Copy the project files to the container
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port your Go application listens on
EXPOSE 8080

# Set metadata labels
LABEL maintainer="ialimoud"
LABEL version="1.0"
LABEL description="ascii-art-web"

# Set the entry point command with the specified port
CMD ["./main", "-port", "8080"]
