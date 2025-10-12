package app

import (
	"database/sql"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humafiber"
	"github.com/gofiber/fiber/v2"
	"github.com/ian-shakespeare/go-app-template/internal/auth"
	"github.com/ian-shakespeare/go-app-template/internal/database"
)

type Request[T any] struct {
	Body T
}

type Response[T any] struct {
	Status int
	Body   T
}

func NewResponse[T any](status int, body T) *Response[T] {
	return &Response[T]{
		Status: status,
		Body:   body,
	}
}

type Empty struct{}

type App struct {
	server *fiber.App
	db     *database.Queries
	op     auth.OAuth2Provider
}

func New(db *sql.DB, op auth.OAuth2Provider) *App {
	server := fiber.New()
	router := humafiber.New(server, huma.DefaultConfig("Go App Template API", "1.0.0"))

	a := &App{
		db: database.New(db),
		op: op,
	}

	api := huma.NewGroup(router, "/api")
	huma.Get(api, "/healthcheck", a.HealthCheck)

	huma.Register(api, huma.Operation{
		Method:      http.MethodPost,
		Path:        "/tasks",
		Summary:     "Create a task",
		Description: "Create a task.",
		Tags:        []string{"Tasks"},
	}, a.CreateTask)

	huma.Register(api, huma.Operation{
		Method:      http.MethodGet,
		Path:        "/tasks",
		Summary:     "List tasks",
		Description: "List all tasks.",
		Tags:        []string{"Tasks"},
	}, a.ListTasks)

	huma.Register(api, huma.Operation{
		Method:      http.MethodPatch,
		Path:        "/tasks/{id}",
		Summary:     "Update a task",
		Description: "Update a task, ignoring empty fields.",
		Tags:        []string{"Tasks"},
	}, a.UpdateTask)

	a.server = server
	return a
}

func (a *App) Listen(addr string) error {
	return a.server.Listen(addr)
}

func (a *App) Test(req *http.Request, msTimeout ...int) (*http.Response, error) {
	return a.server.Test(req, msTimeout...)
}
