.DEFAULT_GOAL := build

# Variables
BINARY_NAME := discord-bot
CMD_DIR := ./cmd/server

# Phony targets
.PHONY: fmt vet build clean run

# Format Go code
fmt:
	@echo "Running go fmt"
	go fmt ./...

# Run go vet for static analysis
vet: fmt
	@echo "Running go vet"
	go vet ./...

# Build the binary
build: vet
	@echo "Running go build"
	go build -o $(BINARY_NAME) $(CMD_DIR)

# Clean up build artifacts
clean:
	@echo "Cleaning up..."
	rm -f $(BINARY_NAME)

# Run the application
run: build
	@echo "Running the application..."
	./$(BINARY_NAME)


