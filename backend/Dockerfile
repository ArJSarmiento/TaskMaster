# Start from the base golang image
FROM golang:latest

# Enable Go Modules
ENV GO111MODULE=on

# Set the Current Working Directory inside the container
WORKDIR /app

# We copy the rest of the source code from the current directory
COPY . .

# Download all the dependencies
RUN go mod download

# Install the package
RUN go install

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["go", "run", "server.go"]