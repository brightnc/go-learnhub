package service

import (
	"fmt"
	"log"
	"os"
	"time"

	redisRepository "github.com/brightnc/go-learnhub/internal/adapter/database/redis/repository"
	"github.com/brightnc/go-learnhub/internal/core/domain"
	"github.com/brightnc/go-learnhub/internal/core/port"
	"github.com/brightnc/go-learnhub/internal/core/util"
	"github.com/golang-jwt/jwt/v5"
)

type UserService struct {
	repo   port.IUserRepository
	blRepo *redisRepository.BlacklistRepository
}

func NewUserService(repo port.IUserRepository, blRepo *redisRepository.BlacklistRepository) *UserService {
	return &UserService{
		repo:   repo,
		blRepo: blRepo,
	}
}

func (svc *UserService) Register(user *domain.User) (*domain.User, error) {
	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword

	user, err = svc.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

type MyCustomClaims struct {
	Id string `json:"id"`
	jwt.RegisteredClaims
}

func (svc *UserService) Login(username string, password string) (string, error) {
	user, err := svc.repo.GetUserByUsername(username)
	if err != nil {
		return "", err
	}

	err = util.ComparePassword(password, user.Password)
	if err != nil {
		log.Fatalf("wrong password : %v", err)
		return "", fmt.Errorf("wrong password")
	}

	JWTSecret := os.Getenv("JWT_SECRET")
	mySigningKey := []byte(JWTSecret)

	claims := MyCustomClaims{
		Id: user.ID.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "learnhub-api",
			Subject:   "user-credential",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		log.Fatalf("Failed to sign token : %v", err)
		return "", fmt.Errorf("failed to sign token")
	}

	return tokenString, nil
}

func (svc *UserService) Logout(token string, expire time.Time) error {
	if err := svc.blRepo.AddToBlackList(token, expire); err != nil {
		return err
	}
	return nil
}

func (svc *UserService) GetUserById(id string) (*domain.User, error) {
	user, err := svc.repo.GetUserById(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
