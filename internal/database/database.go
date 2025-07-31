package database

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"sort"
)

const InitialMigration string = `
create table if not exists migrations (
  id integer primary key,
  version integer not null
);
`

var (
	ErrInvalidMigration    = errors.New("invalid migration")
	ErrMissingMigrationDir = errors.New("migration directory does not exist")
)

func Migrate(db *sql.DB, migrationDir fs.ReadDirFS) error {
	if _, err := db.Exec(InitialMigration); err != nil {
		return fmt.Errorf("%w: failed to create base migration table", ErrInvalidMigration)
	}

	entries, err := migrationDir.ReadDir(".")
	if err != nil {
		return ErrMissingMigrationDir
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name() < entries[j].Name()
	})

	ver := migrationVersion(db)

	i := ver
	for ; i < len(entries); i++ {
		name := entries[i].Name()

		fin, err := migrationDir.Open(name)
		if err != nil {
			return err
		}
		defer fin.Close()

		migration, err := io.ReadAll(fin)
		if err != nil {
			return fmt.Errorf("%w: empty", ErrInvalidMigration)
		}

		if _, err = db.Exec(string(migration)); err != nil {
			return fmt.Errorf("%w: %s", ErrInvalidMigration, name)
		}
	}

	if i == ver {
		return nil
	}

	_, err = db.Exec("insert into migrations (version) values ($1)", i)
	return err
}

func migrationVersion(db *sql.DB) int {
	row := db.QueryRow("select max(version) from migrations")

	var version int
	if err := row.Scan(&version); err != nil {
		return 0
	}

	return version
}
