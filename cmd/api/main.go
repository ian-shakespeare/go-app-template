package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ian-shakespeare/go-app-template/internal/controllers"
	"github.com/ian-shakespeare/go-app-template/internal/models"
)

const DATABASE_FILE string = "example.db"

func main() {
	if err := models.Init(DATABASE_FILE); err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = models.Close()
	}()

	router := controllers.Router()

	addr := ":8080"
	if port, exists := os.LookupEnv("PORT"); exists {
		addr = ":" + port
	}

	fmt.Printf("Starting server on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
