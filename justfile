set windows-shell := ["powershell.exe", "-NoProfile", "-Command"]

tailwind_bin := if os_family() == "windows" { "./tailwindcss.exe" } else { "./tailwindcss" }
binary := if os_family() == "windows" { "main.exe" } else { "main" }

# Default: build then test
default: build test

# Install templ CLI if not present
templ-install:
    if (-not (Get-Command templ -ErrorAction SilentlyContinue)) { Write-Host "Installing templ..."; go install github.com/a-h/templ/cmd/templ@latest }

# Download tailwindcss standalone binary if not present
tailwind-install:
    if (-not (Test-Path "tailwindcss.exe") -and -not (Test-Path "tailwindcss")) { Write-Host "Downloading tailwindcss..."; Invoke-WebRequest -Uri "https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-windows-x64.exe" -OutFile "tailwindcss.exe" }

# Build the binary
build: tailwind-install templ-install
    Write-Host "Building..."
    templ generate
    {{ tailwind_bin }} -i cmd/web/styles/input.css -o cmd/web/assets/css/output.css
    go build -o {{ binary }} cmd/api/main.go

# Run the application
run:
    go run cmd/api/main.go

# Run tests
test:
    go test ./... -v

# Live reload with Air
watch:
    if (-not (Get-Command air -ErrorAction SilentlyContinue)) { Write-Host "Installing air..."; go install github.com/air-verse/air@latest }; air

# Remove build artifacts
clean:
    Remove-Item -ErrorAction SilentlyContinue main, main.exe
