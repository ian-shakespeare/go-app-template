package app

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/ian-shakespeare/go-app-template/internal/database"
)

type CreateTaskRequest struct {
	Name        string     `json:"name"`
	Description *string    `json:"description,omitempty"`
	DueAt       *time.Time `json:"dueAt,omitempty"`
}

type CreateTaskResponse struct {
	Id          int        `json:"id"`
	Name        string     `json:"name"`
	Description *string    `json:"description"`
	DueAt       *time.Time `json:"dueAt"`
	CreatedAt   time.Time  `json:"createdAt"`
}

func (a *App) CreateTask(ctx context.Context, req *Request[CreateTaskRequest]) (*Response[CreateTaskResponse], error) {
	var (
		description sql.NullString
		dueAt       sql.NullTime
	)

	if req.Body.Description != nil {
		description.String = *req.Body.Description
		description.Valid = true
	}

	if req.Body.DueAt != nil {
		dueAt.Time = *req.Body.DueAt
		dueAt.Valid = true
	}

	task, err := a.db.CreateTask(ctx, database.CreateTaskParams{Name: req.Body.Name, Description: description, DueAt: dueAt})
	if err != nil {
		return nil, err
	}

	res := CreateTaskResponse{
		Id:        int(task.TaskID),
		Name:      task.Name,
		CreatedAt: task.CreatedAt,
	}

	if task.Description.Valid {
		res.Description = &task.Description.String
	}

	if task.DueAt.Valid {
		res.DueAt = &task.DueAt.Time
	}

	return NewResponse(http.StatusCreated, res), nil
}
