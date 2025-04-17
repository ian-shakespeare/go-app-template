package models

import (
	"database/sql"
	"os"
	"sort"

	_ "github.com/mattn/go-sqlite3"
)

const MIGRATION_DIR string = "migrations/"

var conn *sql.DB

func Init(connStr string) error {
	db, err := sql.Open("sqlite3", connStr)
	if err != nil {
		return err
	}

	if err := migrate(db); err != nil {
		return err
	}

	conn = db
	return nil
}

func Conn() *sql.DB {
	return conn
}

func Close() error {
	return conn.Close()
}

func migrate(db *sql.DB) error {
	entries, err := os.ReadDir(MIGRATION_DIR)
	if err != nil {
		return err
	}

	// Sort migrations alphanumerically
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name() < entries[j].Name()
	})

	ver := migrationVersion(db)

	i := ver
	for ; i < len(entries); i++ {
		migration, err := os.ReadFile(MIGRATION_DIR + entries[i].Name())
		if err != nil {
			return err
		}

		if _, err := db.Exec(string(migration)); err != nil {
			return err
		}
	}

	if i > ver {
		if _, err := db.Exec("INSERT INTO migrations (version) VALUES (?)", i); err != nil {
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
