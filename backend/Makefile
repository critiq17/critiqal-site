# Makefile for Critiqal Web Site

# Variables

APP_NAME := critiqal-site
MAIN_PACKAGE := ./cmd/main.go
GO := go


# Build the binary

build:
	@echo "Building $(APP_NAME)..."
	$(GO) build -o bin/$(APP_NAME) $(MAIN_PACKAGE)

# Run the app

run:
	@echo "Running $(APP_NAME)..."
	$(GO) run $(MAIN_PACKAGE)

# Clean build artifacts

clean:
	@echo "Cleaning..."
	rm -rf bin/*