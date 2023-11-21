package app

import (
	"html/template"
	"net/http"
	"time"

	"github.com/chris-langager/go-htmx-seed-repo/internal/todo"
)

var todoTemplate = template.Must(template.ParseFS(templateFiles, "todo.html"))

func (o *Server) createTodo(w http.ResponseWriter, r *http.Request) {
	t := todo.Todo{"123", time.Now(), "this is new from the server", false}
	sendHtml(w, todoTemplate, t)
}
