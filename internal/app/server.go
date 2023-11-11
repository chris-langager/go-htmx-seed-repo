package app

import (
	"html/template"
	"net/http"

	"github.com/chris-langager/go-htmx-seed-repo/internal/user"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	router      *chi.Mux
	userService *user.Service
}

// implement http.Gander interface
func (o *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	o.router.ServeHTTP(w, r)
}

/*
Factor function for our server.

It's here that we can handle DI and our routing logic
*/
func NewServer(userService *user.Service) *Server {
	o := &Server{
		router:      chi.NewRouter(),
		userService: userService,
	}

	// put userId on context if a user is logged in
	o.router.Use(userIdMiddleware(userService))

	// static files
	o.router.Handle("/static/*", http.FileServer(http.FS(staticFiles)))

	o.router.Get("/register", o.registerPage)
	o.router.Post("/register", o.register)
	o.router.Get("/login", o.loginPage)
	o.router.Post("/login", o.login)
	o.router.Get("/logout", o.logout)

	o.router.Get("/", o.homePage)

	o.router.NotFound(o.fourOhFourPage)

	return o
}

type ApplicationProperties struct {
	Title string
	User  *user.User
}

func renderPage(w http.ResponseWriter, template *template.Template, data any) {
	err := template.Execute(w, data)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}
}
