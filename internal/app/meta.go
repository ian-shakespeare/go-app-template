package app

import (
	"context"
	"net/http"
)

type HealthCheckResponse struct {
	Status string `json:"status"`
}

func (a *App) HealthCheck(ctx context.Context, _ *EmptyRequest) (*Response[HealthCheckResponse], error) {
	health := Response[HealthCheckResponse]{
		Status: http.StatusOK,
		Body: HealthCheckResponse{
			Status: "OK",
		},
	}

	return &health, nil
}
