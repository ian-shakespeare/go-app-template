package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ian-shakespeare/go-app-template/internal/controllers"
	"github.com/ian-shakespeare/go-app-template/internal/database"
	"github.com/ian-shakespeare/go-app-template/internal/models"
	_ "github.com/mattn/go-sqlite3"
)

const DATABASE_FILE string = "example.db"

func main() {
	db, err := sql.Open("sqlite3", DATABASE_FILE)
	if err != nil {
		log.Fatal(err)
	}

	if err := database.Migrate(context.Background(), db); err != nil {
		log.Fatal(err)
	}

	models.Init(db)

	router := controllers.Router()

	addr := ":8080"
	if port, exists := os.LookupEnv("PORT"); exists {
		addr = ":" + port
	}

	fmt.Printf("Starting server on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
