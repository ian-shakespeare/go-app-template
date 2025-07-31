package database_test

import (
	"database/sql"
	"testing"
	"testing/fstest"

	"github.com/ian-shakespeare/go-app-template/internal/database"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestMigrate(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		db, err := sql.Open("sqlite3", ":memory:")
		if err != nil {
			t.Fatal(err)
		}

		migrationFS := fstest.MapFS{
			"1_table.sql": &fstest.MapFile{
				Data: []byte("create table example (id integer primary key, name text not null);"),
			},
			"2_insert.sql": &fstest.MapFile{
				Data: []byte("insert into example (name) values ('example');"),
			},
		}

		err = database.Migrate(db, migrationFS)
		assert.NoError(t, err)

		var version int
		err = db.QueryRow("select version from migrations").Scan(&version)
		assert.NoError(t, err)
		assert.Equal(t, len(migrationFS), version)
	})

	t.Run("ok existing migration", func(t *testing.T) {
		t.Parallel()

		db, err := sql.Open("sqlite3", ":memory:")
		if err != nil {
			t.Fatal(err)
		}

		migrationFS := fstest.MapFS{
			"1_table.sql": &fstest.MapFile{
				Data: []byte("create table example (id integer primary key, name text not null);"),
			},
		}

		err = database.Migrate(db, migrationFS)
		assert.NoError(t, err)

		var version int
		err = db.QueryRow("select max(version) from migrations").Scan(&version)
		assert.NoError(t, err)
		assert.Equal(t, 1, version)

		migrationFS["2_insert.sql"] = &fstest.MapFile{
			Data: []byte("insert into example (name) values ('example');"),
		}

		err = database.Migrate(db, migrationFS)
		assert.NoError(t, err)

		err = db.QueryRow("select max(version) from migrations").Scan(&version)
		assert.NoError(t, err)
		assert.Equal(t, 2, version)
	})

	t.Run("ok same migration", func(t *testing.T) {
		t.Parallel()

		db, err := sql.Open("sqlite3", ":memory:")
		if err != nil {
			t.Fatal(err)
		}

		migrationFS := fstest.MapFS{
			"1_table.sql": &fstest.MapFile{
				Data: []byte("create table example (id integer primary key, name text not null);"),
			},
			"2_insert.sql": &fstest.MapFile{
				Data: []byte("insert into example (name) values ('example');"),
			},
		}

		err = database.Migrate(db, migrationFS)
		assert.NoError(t, err)

		var version int
		err = db.QueryRow("select max(version) from migrations").Scan(&version)
		assert.NoError(t, err)
		assert.Equal(t, 2, version)

		err = database.Migrate(db, migrationFS)
		assert.NoError(t, err)

		err = db.QueryRow("select max(version) from migrations").Scan(&version)
		assert.NoError(t, err)
		assert.Equal(t, 2, version)
	})

	t.Run("invalid migration", func(t *testing.T) {
		t.Parallel()

		db, err := sql.Open("sqlite3", ":memory:")
		if err != nil {
			t.Fatal(err)
		}

		migrationsFS := fstest.MapFS{
			"1_table.sql": &fstest.MapFile{
				Data: []byte("create bad table"),
			},
		}

		err = database.Migrate(db, migrationsFS)
		assert.Error(t, err)
		assert.ErrorIs(t, err, database.ErrInvalidMigration)
	})
}
