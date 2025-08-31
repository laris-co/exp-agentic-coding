# custom-pb-plus

An improved PocketBase-based Go app derived from `custom-pb` with:

- `/hello`, `/healthz`, and `/version` endpoints
- Lightweight request logging via stdlib `log/slog`
- Build-time version metadata via `-ldflags`
- Basic tests using `httptest`

## Quick Start

```bash
# Build with version metadata
make build VERSION=0.1.0

# Run in dev mode
make run

# Test
make test
```

Or manually:

```bash
go build -o ../bin/custom-pocketbase-plus \
  -ldflags "-X main.version=dev -X main.commit=$(git rev-parse --short HEAD) -X main.builtAt=$(date -u +%Y-%m-%dT%H:%M:%SZ)" \
  .

../bin/custom-pocketbase-plus serve --dev
```

## Endpoints

- `GET /hello` → `Hello world!`
- `GET /healthz` → `{ "status": "ok" }`
- `GET /version` → `{ "version", "commit", "builtAt" }`

## Public Directory

Use PocketBase's `--publicDir` flag to serve static files (defaults to `pb_public`):

```bash
./bin/custom-pocketbase-plus serve --publicDir="./bin/pb_public"
```

## Notes

- Requires Go 1.23+
- PocketBase v0.29.3

