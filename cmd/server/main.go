package main

import (
	"fmt"
	"net/http"

	"github.com/chris-langager/go-htmx-seed-repo/internal/app"
)

func main() {
	app := app.NewServer()

	port := ":3000"

	fmt.Printf("server listening on %s...\n", port)
	err := http.ListenAndServe(port, app)
	if err != nil {
		panic(err)
	}
}
