package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/chris-langager/go-htmx-seed-repo/internal/app"
	"github.com/chris-langager/go-htmx-seed-repo/internal/db"
	"github.com/chris-langager/go-htmx-seed-repo/internal/todo"
	"github.com/chris-langager/go-htmx-seed-repo/internal/user"
)

//TODO: make cfg package for env vars and defaults

func main() {
	postgresUrl := "postgres://postgres:passw0rd@localhost:5432/postgres?sslmode=disable"
	if os.Getenv("DATABASE_URL") != "" {
		postgresUrl = os.Getenv("DATABASE_URL")
	}

	dbConnection, err := sql.Open("postgres", postgresUrl)
	if err != nil {
		panic(err)
	}
	db.MustMigrate(dbConnection)

	userService := user.NewService(dbConnection)
	todoService := todo.NewService((dbConnection))

	app := app.NewServer(userService, todoService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fmt.Printf("server listening on %s...\n", port)
	err = http.ListenAndServe(":"+port, app)
	if err != nil {
		panic(err)
	}
}
