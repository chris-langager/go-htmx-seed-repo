package user

type User struct {
	Id             string `json:"Id"`
	Email          string `json:"email"`
	HashedPassword string `json:"-"`
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
