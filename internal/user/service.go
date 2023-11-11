package user

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

// temp

var users = []User{
	{
		Id:             "123",
		Email:          "test@fake.com",
		HashedPassword: mustHashPassword("test"),
	},
}
