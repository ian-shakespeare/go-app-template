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

	if err := os.WriteFile(MIGRATION_DIR+migrationName, []byte("-- write migration code below"), 0o655); err != nil {
		_, _ = os.Stderr.WriteString("failed to write migration file\n")
		os.Exit(1)
	}

	successMsg := fmt.Sprintf("successfully created migration: %s\n", migrationName)
	_, _ = os.Stdout.WriteString(successMsg)
}
