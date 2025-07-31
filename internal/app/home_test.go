package app_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ian-shakespeare/go-app-template/internal/app"
	"github.com/ian-shakespeare/go-app-template/internal/viewrenderer"
	"github.com/ian-shakespeare/go-app-template/web/templates"
	"github.com/stretchr/testify/assert"
)

func TestHome(t *testing.T) {
	vr, err := viewrenderer.New(templates.FS)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		w := httptest.NewRecorder()
		r := httptest.NewRequestWithContext(t.Context(), "GET", "/home", http.NoBody)
		a := app.New(nil, vr, nil)

		a.Home(w, r)
		res := w.Result()
		defer res.Body.Close()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		page, err := io.ReadAll(res.Body)
		assert.NoError(t, err)
		assert.Contains(t, string(page), "Home page")
	})
}
