package user

import (
	"context"
	"fmt"
)

func (o *Service) GetUser(ctx context.Context, id string) (*User, error) {
	sql := `SELECT * FROM users WHERE id = $1`
	userRows := []userRow{}
	err := o.db.SelectContext(ctx, &userRows, sql, id)
	if err != nil {
		return nil, fmt.Errorf("error looking up user by id: %w", err)
	}

	if len(userRows) == 0 {
		return nil, nil
	}

	return userRows[0].toUser(), nil
}

func (o *Service) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	sql := `SELECT * FROM users WHERE email = $1`
	userRows := []userRow{}
	err := o.db.SelectContext(ctx, &userRows, sql, email)
	if err != nil {
		return nil, fmt.Errorf("error looking up user by email: %w", err)
	}

	if len(userRows) == 0 {
		return nil, nil
	}

	return userRows[0].toUser(), nil
}

func (o *Service) GetUserByLoginInfo(ctx context.Context, email string, password string) (*User, error) {
	user, err := o.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	if !passwordsMatch(user.HashedPassword, password) {
		return nil, nil // Or you could return an error, your call
	}

	return user, nil

}
