package user

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type CreateUserInput struct {
	Email    string
	Password string
}

type userRow struct {
	Id             string `db:"id"`
	Email          string `db:"email"`
	HashedPassword string `db:"hashed_password"`
}

func (o *userRow) toUser() *User {
	return &User{
		Id:             o.Id,
		Email:          o.Email,
		HashedPassword: o.HashedPassword,
	}
}
func (o *Service) CreateUser(ctx context.Context, input CreateUserInput) (*User, error) {
	// TODO - validate, come up with validation error type

	hashedPassword, err := hashPassword(input.Password)
	if err != nil {
		return nil, fmt.Errorf("error hashing hassword: %w", err)
	}

	user := User{
		Id:             uuid.NewString(),
		Email:          input.Email,
		HashedPassword: hashedPassword,
	}

	_, err = o.db.NamedExecContext(ctx, `INSERT INTO users 
						(id, email, hashed_password) 
						VALUES (:id, :email, :hashed_password)`,
		&userRow{user.Id, user.Email, user.HashedPassword})

	if err != nil {
		return nil, fmt.Errorf("error inserting user into db: %w", err)
	}

	return &user, nil
}
