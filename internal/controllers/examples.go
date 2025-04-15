package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/ian-shakespeare/go-app-template/internal/models"
)

func createExample(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		handleError(w, http.StatusBadRequest, "bad request")
		return
	}

	var example models.NewExample
	if err := json.Unmarshal(b, &example); err != nil {
		handleError(w, http.StatusBadRequest, "bad request")
		return
	}

	created, err := models.Examples.Create(example)
	if err != nil {
		handleError(w, commonErrorCodes(err), err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, created)
}

func listExamples(w http.ResponseWriter, r *http.Request) {
	examples, err := models.Examples.All()
	if err != nil {
		handleError(w, commonErrorCodes(err), err.Error())
		return
	}

	writeJSON(w, http.StatusOK, examples)
}

func getExample(w http.ResponseWriter, r *http.Request) {
	exampleIdStr := r.PathValue("exampleId")

	exampleId, err := strconv.Atoi(exampleIdStr)
	if err != nil {
		handleError(w, http.StatusBadRequest, "bad request")
		return
	}

	example, err := models.Examples.Get(exampleId)
	if err != nil {
		handleError(w, commonErrorCodes(err), err.Error())
		return
	}

	writeJSON(w, http.StatusOK, example)
}

func updateExample(w http.ResponseWriter, r *http.Request) {
	exampleIdStr := r.PathValue("exampleId")

	exampleId, err := strconv.Atoi(exampleIdStr)
	if err != nil {
		handleError(w, http.StatusBadRequest, "bad request")
		return
	}

	b, err := io.ReadAll(r.Body)
	if err != nil {
		handleError(w, http.StatusBadRequest, "bad request")
		return
	}

	var example models.EditExample
	if err := json.Unmarshal(b, &example); err != nil {
		handleError(w, http.StatusBadRequest, "bad request")
		return
	}

	updated, err := models.Examples.Update(exampleId, example)
	if err != nil {
		handleError(w, commonErrorCodes(err), err.Error())
		return
	}

	writeJSON(w, http.StatusOK, updated)
}

func deleteExample(w http.ResponseWriter, r *http.Request) {
	exampleIdStr := r.PathValue("exampleId")

	exampleId, err := strconv.Atoi(exampleIdStr)
	if err != nil {
		handleError(w, http.StatusBadRequest, "bad request")
		return
	}

	if err := models.Examples.Delete(exampleId); err != nil {
		handleError(w, commonErrorCodes(err), err.Error())
		return
	}

	writeString(w, http.StatusOK, "deleted")
}
