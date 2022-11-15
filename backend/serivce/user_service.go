package service

import "github.com/axxmk/go-totp-authentication/repository"

type userService struct {
	repository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) userService { // Adapter of Repository
	return userService{repository: userRepository}
}

func (s userService) SignUp(email string, password string) (*string, *string, error) {
	user, err := s.repository.CreateUser(email, password, "")
	if err != nil {
		return nil, nil, err
	}

	return &user.Email, &user.Password, nil
}

func (s userService) SignIn(email string, password string) (*User, error) {
	user, err := s.repository.CheckUser(email)
	if err != nil {
		return nil, err
	}

	payload := &User{
		Id:    user.Id,
		Email: user.Email,
	}

	return payload, nil
}

func (s userService) ListUsers() ([]*User, error) {
	users, err := s.repository.GetUsers()
	if err != nil {
		return nil, err
	}
	var userResponse []*User
	for _, user := range users {
		userResponse = append(userResponse, &User{Id: user.Id, Email: user.Email})
	}
	return userResponse, nil
}
