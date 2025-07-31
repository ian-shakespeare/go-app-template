package viewrenderer

import (
	"html/template"
	"io"
	"io/fs"
	"strings"
)

type ViewRenderer struct {
	tmpl  *template.Template
	views fs.FS
}

func New(templates fs.FS) (*ViewRenderer, error) {
	tmpl, err := template.New("default").ParseFS(templates, "*.html.tmpl")

	return &ViewRenderer{tmpl: tmpl, views: templates}, err
}

func (vr *ViewRenderer) Render(w io.Writer, name string, data any) error {
	if !strings.HasSuffix(name, ".html.tmpl") {
		name += ".html.tmpl"
	}

	tmpl, err := vr.tmpl.Clone()
	if err != nil {
		return err
	}

	tmpl, err = tmpl.ParseFS(vr.views, name)
	if err != nil {
		return err
	}

	return tmpl.ExecuteTemplate(w, name, data)
}
