# Build stage: This stage is used to build the Go application
FROM golang:1.23 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the container
# These files are used to manage Go module dependencies
COPY go.mod go.sum ./

# Download and cache dependencies to avoid re-downloading them on subsequent builds
RUN go mod download

# Copy the entire project source code to the container
COPY . .

# Build the Go application
# CGO_ENABLED=0 disables C bindings, ensuring a statically linked binary
# GOOS=linux specifies that the binary is built for a Linux environment
# The compiled binary will be output as "weather-notification-service"
RUN CGO_ENABLED=0 GOOS=linux go build -o discord-bot ./cmd/server

# Final stage: This stage is used to for running the application in the lightweight runtime image
FROM alpine:latest

# Set the working directory for the final image
WORKDIR /app

# Copy the compiled binary from the builder stage to the final image
COPY --from=builder /app/discord-bot .

# Copy the .env file from the builder stage to the final image
COPY --from=builder /app/.env .

# Ensure the compiled binary has execute permissions
RUN chmod +x discord-bot

# Define the command to run the application when the container starts
CMD ["./discord-bot"]
