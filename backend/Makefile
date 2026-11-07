.PHONY: build run test test-coverage lint

build:
	go build -o bin/server ./cmd/server

run:
	go run ./cmd/server

test:
	go test ./... -v

test-coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report: coverage.html"

lint:
	go vet ./...

swagger:
	swag init -g cmd/server/main.go -o docs

# SDD commands
propose:
	@./scripts/propose.sh $(op)

analyze:
	@./scripts/analyze.sh $(op)

apply:
	@./scripts/apply.sh $(op)

sync:
	@./scripts/sync.sh

verify:
	@./scripts/verify.sh $(op)
