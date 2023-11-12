package todo

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Service struct {
	db *sqlx.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{
		db: sqlx.NewDb(db, "postgres"),
	}
}
