package app

import "net/http"

func (a *App) Home(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var data struct {
		PageTitle       string
		PageDescription string
	}
	data.PageTitle = "Home"
	data.PageDescription = "My home page."

	_ = a.writeTemplate(w, "home", data)
}
