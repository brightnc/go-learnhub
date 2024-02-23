package dto

import (
	"time"

	"github.com/brightnc/go-learnhub/internal/core/domain"
	"github.com/google/uuid"
)

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username"  binding:"required"`
	Password string `json:"password"  binding:"required,min=8"`
	Email    string `json:"email"  binding:"required,email"`
}
type RegisterResponse struct {
	ID        uuid.UUID `json:"id" `
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

type LoginRequest struct {
	Username string `json:"username"  binding:"required"`
	Password string `json:"password"  binding:"required,min=8"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id" `
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}

func ToUserResponse(user domain.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt.Time,
	}
}
