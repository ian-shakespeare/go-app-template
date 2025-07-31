package app

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"

	"github.com/ian-shakespeare/go-app-template/internal/auth"
	"github.com/ian-shakespeare/go-app-template/internal/middleware"
	"github.com/ian-shakespeare/go-app-template/internal/viewrenderer"
)

type AppState string

const (
	Starting AppState = "starting"
	Healthy  AppState = "healthy"
	Degraded AppState = "degraded"
)

type App struct {
	State            AppState
	isLoggingEnabled bool
	router           *http.ServeMux
	db               *sql.DB
	vr               *viewrenderer.ViewRenderer
	op               auth.OAuth2Provider
}

func New(db *sql.DB, vr *viewrenderer.ViewRenderer, op auth.OAuth2Provider) *App {
	h := &App{
		State:  Starting,
		router: http.NewServeMux(),
		db:     db,
		vr:     vr,
		op:     op,
	}

	h.loadHandlers()
	h.State = Healthy
	return h
}

func (a *App) ToggleLogging() {
	a.isLoggingEnabled = !a.isLoggingEnabled
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var h http.Handler
	h = a.router

	if a.isLoggingEnabled {
		h = middleware.NewLogger(h)
	}

	h.ServeHTTP(w, r)
}

func (a *App) loadHandlers() {
	a.router.HandleFunc("GET /api/healthcheck", a.HealthCheck)
	a.router.HandleFunc("GET /example", a.Example)
	a.router.HandleFunc("GET /home", a.Home)
}

func (a *App) writeTemplate(w http.ResponseWriter, name string, data any) error {
	w.Header().Set("Content-Type", "text/html")
	return a.vr.Render(w, name, data)
}

func (a *App) writeJSON(w http.ResponseWriter, v any) error {
	buf := new(bytes.Buffer)
	encoder := json.NewEncoder(buf)
	if err := encoder.Encode(v); err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	_, err := io.Copy(w, buf)
	return err
}
