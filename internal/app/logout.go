package app

import "net/http"

func (o *Server) logout(w http.ResponseWriter, r *http.Request) {
	withUser(r.Context(), nil)
	clearAuthCookies(w)
	o.loginPage(w, r)
}
