BIN := bin/go-app-template
GO := go
LINTER := golangci-lint
GEN_MIGRATION := cmd/genmigration/main.go
WEB := cmd/web/main.go
RUN_FLAGS := -v

all: run

run:
	$(GO) run $(WEB) $(RUN_FLAGS)

build:
	$(GO) build -o $(BIN) $(WEB)


migration:
	$(GO) run $(GEN_MIGRATION) $(MIGRATION_NAME)

test:
	$(GO) test ./internal/... -failfast

lint:
	$(LINTER) run ./...

fmt:
	$(GO) fmt ./...

clean:
	rm -rf $(BIN)

.PHONY: down lint migration test up
