package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
)

func Router() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("POST /examples", createExample)
	router.HandleFunc("GET /examples", listExamples)
	router.HandleFunc("GET /examples/{exampleId}", getExample)
	router.HandleFunc("PUT /examples/{exampleId}", updateExample)
	router.HandleFunc("DELETE /examples/{exampleId}", deleteExample)

	return router
}

func writeJSON(w http.ResponseWriter, statusCode int, v any) {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, _ = w.Write(b)
}

func writeString(w http.ResponseWriter, statusCode int, s string) {
	w.WriteHeader(statusCode)
	_, _ = w.Write([]byte(s))
}

func handleError(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	_, _ = w.Write([]byte(message))
}

func commonErrorCodes(e error) int {
	switch {
	case errors.Is(e, sql.ErrNoRows):
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
