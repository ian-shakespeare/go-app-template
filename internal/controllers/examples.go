package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/ian-shakespeare/go-app-template/internal/models/examples"
)

func createExample(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		handleError(w, http.StatusBadRequest, "bad request")
		return
	}

	var e examples.ExampleNew
	if err := json.Unmarshal(b, &e); err != nil {
		handleError(w, http.StatusBadRequest, "bad request")
		return
	}

	created, err := examples.Create(e)
	if err != nil {
		handleError(w, commonErrorCodes(err), err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, created)
}

func listExamples(w http.ResponseWriter, r *http.Request) {
	examples, err := examples.All()
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

	e, err := examples.Get(exampleId)
	if err != nil {
		handleError(w, commonErrorCodes(err), err.Error())
		return
	}

	writeJSON(w, http.StatusOK, e)
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

	var e examples.ExampleEdit
	if err := json.Unmarshal(b, &e); err != nil {
		handleError(w, http.StatusBadRequest, "bad request")
		return
	}

	updated, err := examples.Update(exampleId, e)
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

	if err := examples.Delete(exampleId); err != nil {
		handleError(w, commonErrorCodes(err), err.Error())
		return
	}

	writeString(w, http.StatusOK, "deleted")
}
