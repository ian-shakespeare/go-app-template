package middleware

import (
	"log/slog"
	"net/http"
)

type Logger struct {
	next http.Handler
}

func NewLogger(next http.Handler) *Logger {
	return &Logger{next}
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mw := newMiddlewareWriter(w)
	l.next.ServeHTTP(mw, r)

	statusCode := mw.StatusCode
	method := r.Method
	path := r.URL.Path

	slog.Info("request", "statusCode", statusCode, "method", method, "path", path)
}
