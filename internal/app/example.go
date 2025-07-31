package app

import "net/http"

func (a *App) Example(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	query := r.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "John Doe"
	}

	var data struct {
		PageTitle       string
		PageDescription string
		Name            string
	}
	data.PageTitle = "Example"
	data.PageDescription = "Example page."
	data.Name = name

	_ = a.writeTemplate(w, "example", data)
}
