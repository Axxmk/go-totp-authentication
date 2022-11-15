package service

type User struct {
	Id    int64  `json:"id" bson:"_id"`
	Email string `json:"email"`
}

type UserService interface {
	SignUp(email string, password string) (*string, *string, error)
	SignIn(email string, password string) (*User, error)
	ListUsers() ([]*User, error)
}
