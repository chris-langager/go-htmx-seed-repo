package app

import (
	"html/template"
	"net/http"
)

var homePageTemplate = template.Must(template.ParseFS(files, "layout.html", "homePage.html"))

func (o *Server) homePage(w http.ResponseWriter, r *http.Request) {
	renderPage(w, homePageTemplate, ApplicationProperties{
		Title: "hHome",
	})
}