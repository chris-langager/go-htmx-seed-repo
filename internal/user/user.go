package user

type User struct {
	Id             string `json:"Id"`
	Email          string `json:"email"`
	HashedPassword string `json:"-"`
}
