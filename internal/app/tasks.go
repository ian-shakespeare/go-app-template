package app

import (
	"context"
	"net/http"
	"time"

	"github.com/ian-shakespeare/go-app-template/internal/database"
)

type CreateTaskRequest struct {
	Name        string     `json:"name"`
	Description *string    `json:"description,omitempty"`
	DueAt       *time.Time `json:"due_at,omitempty"`
}

type CreateTaskResponse database.Task

func (a *App) CreateTask(ctx context.Context, req *Request[CreateTaskRequest]) (*Response[CreateTaskResponse], error) {
	task, err := a.db.CreateTask(ctx, database.CreateTaskParams{
		Name:        req.Body.Name,
		Description: req.Body.Description,
		DueAt:       req.Body.DueAt,
	})
	if err != nil {
		return nil, err
	}

	return NewResponse(http.StatusCreated, CreateTaskResponse(task)), nil
}

type ListTasksResponse []database.Task

func (a *App) ListTasks(ctx context.Context, _ *Empty) (*Response[ListTasksResponse], error) {
	tasks, err := a.db.ListTasks(ctx)
	if err != nil {
		return nil, err
	}

	if tasks == nil {
		tasks = []database.Task{}
	}

	return NewResponse(http.StatusOK, ListTasksResponse(tasks)), nil
}

type UpdateTaskRequest struct {
	Id   int64 `path:"id" doc:"ID of the task to update"`
	Body struct {
		Name        string     `json:"name"`
		Description *string    `json:"description"`
		DueAt       *time.Time `json:"due_at"`
	}
}

type UpdateTaskResponse database.Task

func (a *App) UpdateTask(ctx context.Context, req *UpdateTaskRequest) (*Response[UpdateTaskResponse], error) {
	task, err := a.db.UpdateTask(ctx, database.UpdateTaskParams{
		TaskID:      req.Id,
		Name:        req.Body.Name,
		Description: req.Body.Description,
		DueAt:       req.Body.DueAt,
	})
	if err != nil {
		return nil, err
	}

	return NewResponse(http.StatusOK, UpdateTaskResponse(task)), nil
}
