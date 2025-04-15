GO := go
API_MAIN := cmd/api/main.go
SQL_MAIN := cmd/sql/main.go
BIN := bin/

all: run

run:
	$(GO) run $(API_MAIN)

build:
	$(GO) build -o $(BIN)api $(API_MAIN)

migration:
	$(GO) run $(SQL_MAIN) $(MIGRATION_NAME)

lint:
	golangci-lint run ./...

clean:
	rm -rf $(BIN) *.db

.PHONY: all run lint
