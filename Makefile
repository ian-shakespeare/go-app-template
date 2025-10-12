BIN := bin/go-app-template
GO := go
LINTER := golangci-lint
API := cmd/api/main.go
DOCS := cmd/docs/main.go
RUN_FLAGS := -v

all: run

build:
	$(GO) build -o $(BIN) $(API)

run:
	$(GO) run $(API) $(RUN_FLAGS)

docs:
	$(GO) run $(DOCS) docs/openapi.yaml

test:
	$(GO) test ./internal/... -failfast

lint:
	$(LINTER) run ./...

fmt:
	$(GO) fmt ./...

clean:
	rm -rf $(BIN)

.PHONY: run docs lint test up
