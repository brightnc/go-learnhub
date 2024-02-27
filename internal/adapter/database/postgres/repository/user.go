package repository

import (
	"errors"

	"github.com/brightnc/go-learnhub/internal/core/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(user *domain.User) (*domain.User, error) {
	result := r.db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("cannot create user")
	}

	return &domain.User{
		ID:        user.ID,
		Name:      user.Name,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
	}, nil
}

func (r *UserRepository) GetUserByUsername(username string) (*domain.User, error) {
	var user domain.User
	result := r.db.Where("username=?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &domain.User{
		ID:        user.ID,
		Name:      user.Name,
		Username:  user.Username,
		Password:  user.Password,
		DeletedAt: user.DeletedAt,
	}, nil
}

func (r *UserRepository) GetUserById(id string) (*domain.User, error) {
	var user domain.User
	result := r.db.First(&user, "id=?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &domain.User{
		ID:        user.ID,
		Name:      user.Name,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}, nil
}
