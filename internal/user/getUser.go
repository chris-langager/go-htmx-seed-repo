package user

import "context"

func (o *Service) GetUser(ctx context.Context, id string) (*User, error) {
	for _, user := range users {
		if user.Id == id {
			return &user, nil
		}
	}

	return nil, nil
}

func (o *Service) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	for _, user := range users {
		if user.Email == email {
			return &user, nil
		}
	}

	return nil, nil
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
