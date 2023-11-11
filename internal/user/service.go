package user

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Service struct {
	db *sqlx.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{
		db: sqlx.NewDb(db, "postgres"),
	}
}

// temp

var users = []User{
	{
		Id:             "123",
		Email:          "test@fake.com",
		HashedPassword: mustHashPassword("test"),
	},
}
