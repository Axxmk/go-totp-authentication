package handler

import (
	"encoding/json"
	service "github.com/axxmk/go-totp-authentication/serivce"
	"net/http"
)

type userHandler struct {
	service service.UserService
}

func NewUserHandler(userService service.UserService) userHandler { // Adapter of Service
	return userHandler{service: userService}
}

func (h userHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	email, password, err := h.service.SignUp("", "")
	if err != nil {
		return
	}

	response, err := json.Marshal(map[string]any{
		"email":    email,
		"password": password,
	})

	w.Write(response)
}

func (h userHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	user, err := h.service.SignIn("", "")
	if err != nil {
		return
	}

	response, err := json.Marshal(map[string]any{
		"user": user,
	})

	w.Write(response)
}

func (h userHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	// implement me
	h.service.ListUsers()
}
