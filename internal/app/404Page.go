package app

import (
	"html/template"
	"net/http"
)

var fourOhFourPageTemplate = template.Must(template.ParseFS(templateFiles, "_layout.html", "404Page.html"))

func (o *Server) fourOhFourPage(w http.ResponseWriter, r *http.Request) {
	renderPage(w, fourOhFourPageTemplate, ApplicationProperties{
		Title: "Not Found",
		User:  getUser(r.Context()),
	})
}
