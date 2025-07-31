package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/ian-shakespeare/go-app-template/internal/app"
	"github.com/ian-shakespeare/go-app-template/internal/auth"
	"github.com/ian-shakespeare/go-app-template/internal/database"
	"github.com/ian-shakespeare/go-app-template/internal/env"
	"github.com/ian-shakespeare/go-app-template/internal/viewrenderer"
	"github.com/ian-shakespeare/go-app-template/migrations"
	"github.com/ian-shakespeare/go-app-template/web/templates"
	_ "github.com/mattn/go-sqlite3"
)

const (
	dataDir = "data"
)

func loadDotEnv() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		name := entry.Name()
		if !strings.HasSuffix(name, ".env") {
			continue
		}

		fin, err := os.OpenFile(name, os.O_RDONLY, os.ModePerm)
		if err != nil {
			return err
		}
		defer fin.Close()

		if err := env.Load(fin); err != nil {
			return fmt.Errorf("%s: %w", name, err)
		}
	}

	return nil
}

func setupDirectories(baseDir string) error {
	dirs := []string{
		baseDir,
		filepath.Join(baseDir, dataDir),
	}

	for _, dir := range dirs {
		if err := createDirIfNotExists(dir); err != nil {
			return err
		}
	}

	return nil
}

func createDirIfNotExists(path string) error {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		readWriteExecute := os.FileMode(0755)
		err := os.MkdirAll(path, readWriteExecute)
		return err
	}
	if err != nil {
		return err
	}

	if !info.IsDir() {
		return fmt.Errorf("path exists but is not a directory: %s", path)
	}

	return nil
}

func checkDirPermission(path string) error {
	testFile := filepath.Join(path, ".write_test")

	file, err := os.Create(testFile)
	if err != nil {
		return err
	}
	file.Close()

	err = os.Remove(testFile)
	return err
}

func main() {
	verbosePtr := flag.Bool("v", false, "verbose")
	flag.Parse()

	_ = loadDotEnv()

	baseDir := env.Fallback("BASE_DIR", "/var/lib/go-app-template")

	if err := setupDirectories(baseDir); err != nil {
		log.Fatal(err)
	}

	if err := checkDirPermission(baseDir); err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("sqlite3", filepath.Join(baseDir, dataDir, "go-app-template.db"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := database.Migrate(db, migrations.FS); err != nil {
		log.Fatal(err)
	}

	vr, err := viewrenderer.New(templates.FS)
	if err != nil {
		log.Fatal(err)
	}

	op := auth.NewGoogleOAuth2(env.Must(env.Get("GOOGLE_CLIENT_ID")), env.Must(env.Get("GOOGLE_CLIENT_SECRET")))

	a := app.New(db, vr, op)

	if verbosePtr != nil && *verbosePtr {
		a.ToggleLogging()
	}

	addr := ":8000"
	fmt.Printf("Listening on %s\n", addr)
	if err := http.ListenAndServe(addr, a); err != nil {
		log.Fatal(err)
	}
}
