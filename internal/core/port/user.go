package port

import (
	"time"

	"github.com/brightnc/go-learnhub/internal/core/domain"
)

type IUserRepository interface {
	CreateUser(user *domain.User) (*domain.User, error)
	GetUserById(id string) (*domain.User, error)
	GetUserByUsername(username string) (*domain.User, error)
}

type IUserService interface {
	Register(user *domain.User) (*domain.User, error)
	GetUserById(id string) (*domain.User, error)
	Login(username string, password string) (string, error)
	Logout(token string, expire time.Time) error
}
