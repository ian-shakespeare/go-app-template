package app

import (
	"context"
	"database/sql"
	"time"

	"github.com/ian-shakespeare/go-app-template/internal/database"
)

type CreateTaskRequest struct {
	Name        string     `json:"name"`
	Description *string    `json:"description"`
	DueAt       *time.Time `json:"dueAt"`
}

type CreateTaskResponse struct {
	Id          int        `json:"id"`
	Name        string     `json:"name"`
	Description *string    `json:"description"`
	DueAt       *time.Time `json:"dueAt"`
	CreatedAt   time.Time  `json:"createdAt"`
}

func (a *App) CreateTask(ctx context.Context, req *CreateTaskRequest) (*CreateTaskResponse, error) {
	var (
		description sql.NullString
		dueAt       sql.NullTime
	)

	if req.Description != nil {
		description.String = *req.Description
		description.Valid = true
	}

	if req.DueAt != nil {
		dueAt.Time = *req.DueAt
		dueAt.Valid = true
	}

	task, err := a.db.CreateTask(ctx, database.CreateTaskParams{Name: req.Name, Description: description, DueAt: dueAt})
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

	return &res, nil
}
