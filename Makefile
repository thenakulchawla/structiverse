# Makefile

# Default target
.PHONY: all
all: test

# Target for testing
.PHONY: test
test:
	go test ./...

.PHONY: lint
lint:
	revive ./...
