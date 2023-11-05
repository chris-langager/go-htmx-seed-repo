package app

import (
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	router *chi.Mux
}

// implement http.Gander interface
func (o *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	o.router.ServeHTTP(w, r)
}

/*
Factor function for our server.

It's here that we can handle DI and our routing logic
*/
func NewServer() *Server {
	o := &Server{
		router: chi.NewRouter(),
	}

	o.router.Get("/", o.homePage)

	o.router.Get("/*", o.fourOhFourPage)

	return o
}

type ApplicationProperties struct {
	Title string
}

func renderPage(w http.ResponseWriter, template *template.Template, data any) {
	err := template.Execute(w, data)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}
}
