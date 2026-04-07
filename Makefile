# Simple Makefile for a Go project

# Build the application
all: build test

templ-install:
	@command -v templ > /dev/null 2>&1 || { \
		echo "Installing templ..."; \
		go install github.com/a-h/templ/cmd/templ@latest; \
	}

tailwind-install:
	@if [ ! -f tailwindcss ] && [ ! -f tailwindcss.exe ]; then \
		echo "Downloading tailwindcss..."; \
		if [ "$(OS)" = "Windows_NT" ]; then \
			curl -sLo tailwindcss.exe https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-windows-x64.exe; \
		else \
			curl -sLo tailwindcss https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64; \
			chmod +x tailwindcss; \
		fi; \
	fi

TAILWIND_BIN := $(if $(filter Windows_NT,$(OS)),./tailwindcss.exe,./tailwindcss)
BINARY := $(if $(filter Windows_NT,$(OS)),main.exe,main)

build: tailwind-install templ-install
	@echo "Building..."
	@templ generate
	@$(TAILWIND_BIN) -i cmd/web/styles/input.css -o cmd/web/assets/css/output.css
	@go build -o $(BINARY) cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go

# Test the application
test:
	@echo "Testing..."
	@go test ./... -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main main.exe

# Live Reload
watch:
	@command -v air > /dev/null 2>&1 || { echo "Installing air..."; go install github.com/air-verse/air@latest; }; air

.PHONY: all build run test clean watch tailwind-install templ-install
