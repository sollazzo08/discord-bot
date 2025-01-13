.DEFAULT_GOAL := build

.PHONY: fmt vet build clean
fmt:
	@echo "Running go fmt"
	go fmt ./...
vet: fmt
	@echo "Running go vet"
	go vet ./...
build: vet
	@echo "Running go build"
	go build
clean:
	@echo "Cleaning up..."
	rm -f discord-bot
