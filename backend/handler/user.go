package handler

import (
	"encoding/json"
	service "github.com/axxmk/go-totp-authentication/serivce"
	"github.com/axxmk/go-totp-authentication/types"
	"github.com/axxmk/go-totp-authentication/utils"
	"net/http"
)

type userHandler struct {
	service service.UserService
}

func NewUserHandler(userService service.UserService) userHandler { // Adapter of Service
	return userHandler{service: userService}
}

func (h userHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Set the response header to application/json
	w.Header().Set("Content-Type", "application/json")
	var response []byte
	var body types.SignUp
	err := utils.Parse(r, &body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Call signup service
	token, base64, err := h.service.SignUp(body.Email, body.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a response
	response, _ = json.Marshal(map[string]any{"success": true, "token": token, "image": base64})
	w.Write(response)

	return
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
