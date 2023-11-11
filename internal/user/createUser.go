package user

import (
	"context"

	"github.com/google/uuid"
)

type CreateUserInput struct {
	Email    string
	Password string
}

func (o *Service) CreateUser(ctx context.Context, input CreateUserInput) (*User, error) {
	// TODO - validate

	hashedPassword, err := hashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	user := User{
		Id:             uuid.NewString(),
		Email:          input.Email,
		HashedPassword: hashedPassword,
	}

	// TODO - persist to DB
	users = append(users, user)

	return &user, nil
}
