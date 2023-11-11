package app

import (
	"html/template"
	"net/http"

	"github.com/chris-langager/go-htmx-seed-repo/internal/user"
)

var registerPageTemplate = template.Must(template.ParseFS(templateFiles, "_layout.html", "registerPage.html"))

type regiserPageProperties struct {
	ApplicationProperties
	Errors map[string]string
}

func (o *Server) registerPage(w http.ResponseWriter, r *http.Request) {
	if getUser(r.Context()) != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	properties := newRegisterPageProperties()

	renderPage(w, registerPageTemplate, properties)
}

func (o *Server) register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	properties := newRegisterPageProperties()

	email := r.FormValue("email")
	password := r.FormValue("password")

	if email == "" {
		properties.Errors["Email"] = "email is required"
	}

	if password == "" {
		properties.Errors["Password"] = "password is required"
	}

	if len(properties.Errors) > 0 {
		renderPage(w, registerPageTemplate, properties)
		return
	}

	user, err := o.userService.CreateUser(ctx, user.CreateUserInput{
		Email:    email,
		Password: password,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	setAuthCookies(w, user)
	http.Redirect(w, r, "/", http.StatusFound)
}

func newRegisterPageProperties() regiserPageProperties {
	return regiserPageProperties{
		ApplicationProperties: ApplicationProperties{
			Title: "Register",
		},
		Errors: map[string]string{},
	}
}
