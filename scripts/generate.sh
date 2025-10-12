#!/usr/bin/env bash

OPENAPI_GENERATOR="cmd/docs/main.go"
OPENAPI_OUTPUT="docs/openapi.yaml"
TYPES_OUTPUT="web/src/lib/schema.d.ts"

sqlc generate
go run $OPENAPI_GENERATOR $OPENAPI_OUTPUT
npx openapi-typescript $OPENAPI_OUTPUT -o $TYPES_OUTPUT
