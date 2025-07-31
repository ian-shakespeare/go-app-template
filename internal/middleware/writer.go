package middleware

import "net/http"

type middlewareWriter struct {
	http.ResponseWriter

	StatusCode int
}

func newMiddlewareWriter(w http.ResponseWriter) *middlewareWriter {
	return &middlewareWriter{
		ResponseWriter: w,
		StatusCode:     http.StatusOK,
	}
}

func (m *middlewareWriter) WriteHeader(statusCode int) {
	m.ResponseWriter.WriteHeader(statusCode)
	m.StatusCode = statusCode
}
