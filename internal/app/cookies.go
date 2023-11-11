package app

import (
	"context"
	"net/http"
	"time"

	"github.com/chris-langager/go-htmx-seed-repo/internal/user"
)

var userIdCookieName = "userId"

func setAuthCookies(w http.ResponseWriter, user *user.User) {
	http.SetCookie(w, &http.Cookie{
		Name:    userIdCookieName,
		Value:   user.Id,
		Path:    "/",
		Expires: time.Now().Add(365 * 24 * time.Hour),
	})
}

func clearAuthCookies(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:    userIdCookieName,
		Value:   "",
		Path:    "/",
		Expires: time.Now(),
	})
}

func userIdMiddleware(userService *user.Service) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userIdCookie, _ := r.Cookie(userIdCookieName)

			if userIdCookie == nil || userIdCookie.Value == "" {
				next.ServeHTTP(w, r)
				return
			}

			// TODO: swap out userIDCookie and actually validate beause this isn't real auth

			ctx := r.Context()
			user, err := userService.GetUser(ctx, userIdCookie.Value)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}

			if user == nil {
				clearAuthCookies(w)
				next.ServeHTTP(w, r)
			}

			ctx = withUser(ctx, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func authMiddlware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/login" {
			next.ServeHTTP(w, r)
			return
		}

		if getUser(r.Context()) == nil {
			http.Redirect(w, r, "/login", 302)
			return
		}

		next.ServeHTTP(w, r)
	})
}

type key int

const (
	userIdKey key = iota
)

func withUser(ctx context.Context, user *user.User) context.Context {
	return context.WithValue(ctx, userIdKey, user)
}

func getUser(ctx context.Context) *user.User {
	user, ok := ctx.Value(userIdKey).(*user.User)
	if !ok {
		return nil
	}
	return user
}
