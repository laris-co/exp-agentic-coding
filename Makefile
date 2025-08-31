SHELL := /bin/bash

VERSION ?= 0.1.0
COMMIT  ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo dev)
BUILT_AT?= $(shell date -u +%Y-%m-%dT%H:%M:%SZ)

.PHONY: build run test lint

build:
	GOFLAGS= CGO_ENABLED=0 go build -o bin/custom-pocketbase-plus \
		-ldflags "-X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.builtAt=$(BUILT_AT)" \
		./custom-pb-plus

run: build
	./bin/custom-pocketbase-plus serve --dev

test:
	go test ./custom-pb-plus -v

lint:
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run ./custom-pb-plus/... ; \
	else \
		echo "golangci-lint not found; running go vet as fallback"; \
		go vet ./custom-pb-plus/... ; \
	fi

