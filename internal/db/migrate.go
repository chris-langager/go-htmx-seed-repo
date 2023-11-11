package db

import (
	"database/sql"
	"embed"
	"errors"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/lib/pq"
)

var (
	//go:embed migrations/*.sql
	MigrationFiles embed.FS
)

func MustMigrate(db *sql.DB) {
	sourceDriver, err := iofs.New(MigrationFiles, "migrations")
	if err != nil {
		panic(err)
	}

	databaseDriver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		panic(err)
	}

	migrator, err := migrate.NewWithInstance("iofs", sourceDriver, "postgres", databaseDriver)
	if err != nil {
		panic(err)
	}

	err = migrator.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		panic(err)
	}
}
