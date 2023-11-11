package app

import (
	"html/template"
	"net/http"
)

var loginPageTemplate = template.Must(template.ParseFS(templateFiles, "_layout.html", "loginPage.html"))

type loginPageProperties struct {
	ApplicationProperties
	Errors map[string]string
}

func (o *Server) loginPage(w http.ResponseWriter, r *http.Request) {
	if getUser(r.Context()) != nil {
		http.Redirect(w, r, "/", 302)
		return
	}

	properties := newLoginPageProperties()

	renderPage(w, loginPageTemplate, properties)
}

func (o *Server) login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	properties := newLoginPageProperties()

	email := r.FormValue("email")
	password := r.FormValue("password")

	if email == "" {
		properties.Errors["Email"] = "email is required"
	}

	if password == "" {
		properties.Errors["Password"] = "password is required"
	}

	if len(properties.Errors) > 0 {
		renderPage(w, loginPageTemplate, properties)
		return
	}

	user, err := o.userService.GetUserByLoginInfo(ctx, email, password)
	if err != nil {
		//TODO error page
		panic(err)
	}

	if user == nil {
		properties.Errors["Password"] = "incorrect password"
		renderPage(w, loginPageTemplate, properties)
		return
	}

	setAuthCookies(w, user)
	http.Redirect(w, r, "/", 302)
}

func newLoginPageProperties() loginPageProperties {
	return loginPageProperties{
		ApplicationProperties: ApplicationProperties{
			Title: "Login",
		},
		Errors: map[string]string{},
	}
}
