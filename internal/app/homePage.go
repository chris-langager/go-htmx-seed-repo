package app

import (
	"html/template"
	"net/http"

	"github.com/chris-langager/go-htmx-seed-repo/internal/todo"
)

var homePageTemplate = template.Must(template.ParseFS(templateFiles, "_layout.html", "homePage.html", "todo.html"))

type HomePageProperties struct {
	ApplicationProperties
	Todos []todo.Todo
}

func (o *Server) homePage(w http.ResponseWriter, r *http.Request) {
	todos, err := o.todoService.ListTodos(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	renderPage(w, homePageTemplate, HomePageProperties{
		ApplicationProperties: ApplicationProperties{
			Title: "Home",
			User:  getUser(r.Context()),
		},
		Todos: todos,
	})
}
