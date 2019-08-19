package user

type User struct {
	Id       uint    `json:"id"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Bio      string  `json:"bio"`
	Image    *string `json:"image"`
	Password string  `json:"password"`
	Token    *string `json:"token"`
}

type UserResponse struct {
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Bio      string  `json:"bio"`
	Image    *string `json:"image"`
	Token    *string `json:"token"`
}

type Request struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}

type Response struct {
	User   UserResponse `json:"user"`
	Status string       `json:"status"`
}

var Users = []User{
	User{1, "Jake", "jake@test.com", "I work at statefarm", nil, "jake1234", nil},
	User{2, "Jacob", "jacob@test.com", "I work at home", nil, "jacob1234", nil},
}
