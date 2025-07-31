package app_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ian-shakespeare/go-app-template/internal/app"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	states := []app.AppState{
		app.Starting,
		app.Healthy,
		app.Degraded,
	}

	for _, state := range states {
		stateStr := string(state)
		testName := fmt.Sprintf("ok %s", strings.ToLower(stateStr))

		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			w := httptest.NewRecorder()
			r := httptest.NewRequestWithContext(t.Context(), "GET", "/api/healthcheck", http.NoBody)
			h := app.New(nil, nil, nil)
			h.State = state

			h.HealthCheck(w, r)
			res := w.Result()
			defer res.Body.Close()

			var body app.HealthCheckResponse
			decoder := json.NewDecoder(res.Body)
			err := decoder.Decode(&body)

			assert.Equal(t, http.StatusOK, res.StatusCode)
			assert.NoError(t, err)
			assert.Equal(t, stateStr, body.Status)
		})
	}
}
