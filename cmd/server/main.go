package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/chris-langager/go-htmx-seed-repo/internal/app"
	"github.com/chris-langager/go-htmx-seed-repo/internal/db"
	"github.com/chris-langager/go-htmx-seed-repo/internal/user"
)

func main() {

	postgresUrl := "postgres://postgres:passw0rd@localhost:5432/postgres?sslmode=disable"
	if os.Getenv("DATABASE_PRIVATE_URL") != "" {
		postgresUrl = os.ExpandEnv("DATABASE_PRIVATE_URL")
	}

	fmt.Printf("postgresUrl='%s'\n", postgresUrl)

	dbConnection, err := sql.Open("postgres", postgresUrl)
	if err != nil {
		panic(err)
	}
	db.MustMigrate(dbConnection)

	app := app.NewServer(user.NewService(dbConnection))

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
