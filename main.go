package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/chris-langager/go-htmx-seed-repo/internal/app"
	"github.com/chris-langager/go-htmx-seed-repo/internal/user"
)

func main() {
	app := app.NewServer(user.NewService())

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fmt.Printf("server listening on %s...\n", port)
	err := http.ListenAndServe(":"+port, app)
	if err != nil {
		panic(err)
	}
}
