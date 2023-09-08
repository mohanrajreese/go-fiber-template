# Use the official GoLang base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /sample-server

# Copy the Go module files
COPY go.mod go.sum ./

# Download and install project dependencies
RUN go mod download

# Copy the rest of the project files
COPY . .

# Build the Go application
RUN go build -o sample-gofiber-server

# Set the entry point command
CMD ["./sample-gofiber-server"]

# Expose the port on which the application listens
EXPOSE 2105