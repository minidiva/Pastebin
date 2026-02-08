package handler

import (
	"auth/internal/entity"
	"encoding/json"
	"fmt"
	"net/http"
)

type UserService interface {
	CreateUser(user entity.User) error
}

type UserHandler struct {
	service UserService
}

func NewUserHandler(s UserService) *UserHandler {
	return &UserHandler{
		service: s,
	}
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// Method Check
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)

		fmt.Fprintf(w, "wrong method")
	}

	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	user := entity.User{
		Username: req.Username,
		Email:    req.Email,
		// Password будет захеширован в service слое
	}

	err := h.service.CreateUser(user, req.Password)

	return
}

func (h *UserHandler) HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func (h *UserHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
}

// ID           int       `db:"id"`
// Username     string    `db:"username"`
// PasswordHash string    `db:"password_hash"`
// Email        string    `db:"email"`
// CreatedAt    time.Time `db:"created_at"`
