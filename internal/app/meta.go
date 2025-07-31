package app

import "net/http"

type HealthCheckResponse struct {
	Status string `json:"status"`
}

func (a *App) HealthCheck(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var hcr HealthCheckResponse
	hcr.Status = string(a.State)

	_ = a.writeJSON(w, hcr)
}
