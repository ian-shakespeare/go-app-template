package database

import (
	"context"
	"database/sql"
	"os"
	"sort"
)

const MIGRATION_DIR string = "migrations/"

func Migrate(ctx context.Context, db *sql.DB) error {
	entries, err := os.ReadDir(MIGRATION_DIR)
	if err != nil {
		return err
	}

	// Sort migrations alphanumerically
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name() > entries[j].Name()
	})

	ver := migrationVersion(db)

	i := ver
	for ; i < len(entries); i++ {
		migration, err := os.ReadFile(MIGRATION_DIR + entries[i].Name())
		if err != nil {
			return err
		}

		if _, err := db.ExecContext(ctx, string(migration)); err != nil {
			return err
		}
	}

	if i > ver {
		if _, err := db.ExecContext(ctx, "INSERT INTO migrations (version) VALUES (?)", i); err != nil {
			return err
		}
	}

	return nil
}

func migrationVersion(db *sql.DB) int {
	row := db.QueryRow("SELECT MAX(version) FROM migrations")

	var version int
	if err := row.Scan(&version); err != nil {
		return 0
	}

	return version
}
