package viewrenderer_test

import (
	"bytes"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/ian-shakespeare/go-app-template/internal/viewrenderer"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		_, err := viewrenderer.New(fstest.MapFS{
			"base.html.tmpl": &fstest.MapFile{
				Data: []byte("base"),
			},
		})
		assert.NoError(t, err)
	})

	t.Run("empty", func(t *testing.T) {
		t.Parallel()

		_, err := viewrenderer.New(fstest.MapFS{})
		assert.Error(t, err)
	})
}

func TestRender(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		vr, err := viewrenderer.New(fstest.MapFS{
			"base.html.tmpl": &fstest.MapFile{
				Data: []byte(`base`),
			},
		})
		if err != nil {
			t.Fatal(err)
		}

		buf := new(bytes.Buffer)
		err = vr.Render(buf, "base", nil)
		assert.NoError(t, err)
		assert.Equal(t, "base", strings.Trim(buf.String(), " \r\n"))
	})

	t.Run("ok override", func(t *testing.T) {
		t.Parallel()

		vr, err := viewrenderer.New(fstest.MapFS{
			"base.html.tmpl": &fstest.MapFile{
				Data: []byte(`header {{ block "content" . }} {{ end }} footer`),
			},
			"override.html.tmpl": &fstest.MapFile{
				Data: []byte(`{{ template "base.html.tmpl" . }} {{ define "content" }}override{{ end }}`),
			},
		})
		if err != nil {
			t.Fatal(err)
		}

		buf := new(bytes.Buffer)
		err = vr.Render(buf, "override", nil)
		assert.NoError(t, err)
		assert.Equal(t, "header override footer", strings.Trim(buf.String(), " \r\n"))
	})

	t.Run("ok multiple overrides", func(t *testing.T) {
		t.Parallel()

		vr, err := viewrenderer.New(fstest.MapFS{
			"base.html.tmpl": &fstest.MapFile{
				Data: []byte(`header {{ block "content" . }} {{ end }} footer`),
			},
			"override1.html.tmpl": &fstest.MapFile{
				Data: []byte(`{{ template "base.html.tmpl" . }} {{ define "content" }}override1{{ end }}`),
			},
			"override2.html.tmpl": &fstest.MapFile{
				Data: []byte(`{{ template "base.html.tmpl" . }} {{ define "content" }}override2{{ end }}`),
			},
		})
		if err != nil {
			t.Fatal(err)
		}

		buf := new(bytes.Buffer)
		err = vr.Render(buf, "override1", nil)
		assert.NoError(t, err)
		assert.Equal(t, "header override1 footer", strings.Trim(buf.String(), " \r\n"))

		buf.Reset()

		err = vr.Render(buf, "override2", nil)
		assert.NoError(t, err)
		assert.Equal(t, "header override2 footer", strings.Trim(buf.String(), " \r\n"))
	})

	t.Run("missing", func(t *testing.T) {
		t.Parallel()

		vr, err := viewrenderer.New(fstest.MapFS{
			"base.html.tmpl": &fstest.MapFile{
				Data: []byte(`base`),
			},
		})
		if err != nil {
			t.Fatal(err)
		}

		buf := new(bytes.Buffer)
		err = vr.Render(buf, "bad", nil)
		assert.Error(t, err)
	})
}
