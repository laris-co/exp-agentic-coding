# Custom PocketBase Application

A custom Go application using PocketBase as a framework with additional routes.

## Features

- All standard PocketBase functionality
- Custom `/hello` endpoint

## Quick Start

```bash
# Build
go build -o custom-pocketbase

# Run
./custom-pocketbase serve

# Test custom endpoint
curl http://127.0.0.1:8090/hello
```

## Endpoints

- `/hello` - Returns "Hello world!"
- `/_/` - PocketBase Admin UI
- `/api/` - PocketBase REST API

## Development

```bash
# Run in development mode
go run main.go serve --dev

# Run on different port
go run main.go serve --http="127.0.0.1:8091"
```