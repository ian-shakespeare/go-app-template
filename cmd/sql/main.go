package main

import (
	"fmt"
	"os"
	"time"
)

const MIGRATION_DIR string = "migrations/"

func main() {
	if len(os.Args) < 2 {
		_, _ = os.Stderr.WriteString("usage: command [migration name]\n")
		os.Exit(1)
	}

	migrationName := fmt.Sprintf("%d_%s.sql", time.Now().Unix(), os.Args[1])

	if err := os.WriteFile(MIGRATION_DIR+migrationName, []byte("-- Write migration code below"), 0655); err != nil {
		_, _ = os.Stderr.WriteString("failed to create migration file\n")
		os.Exit(1)
	}

	successMessage := fmt.Sprintf("Successfully created migration: %s\n", migrationName)
	_, _ = os.Stdout.WriteString(successMessage)
}
