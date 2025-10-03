BIN := bin/go-app-template
GO := go
LINTER := golangci-lint
MAIN := cmd/main.go
RUN_FLAGS := -v

all: run

build:
	$(GO) build -o $(BIN) $(MAIN)

run:
	$(GO) run $(MAIN) $(RUN_FLAGS)

test:
	$(GO) test ./internal/... -failfast

lint:
	$(LINTER) run ./...

fmt:
	$(GO) fmt ./...

clean:
	rm -rf $(BIN)

.PHONY: run lint test up
