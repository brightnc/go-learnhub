package protocol

import (
	"fmt"

	"github.com/brightnc/go-learnhub/internal/adapter/database/postgres"
	"github.com/brightnc/go-learnhub/internal/adapter/database/postgres/repository"
	"github.com/brightnc/go-learnhub/internal/adapter/database/redis"
	redisRepository "github.com/brightnc/go-learnhub/internal/adapter/database/redis/repository"
	"github.com/brightnc/go-learnhub/internal/core/service"
	"github.com/joho/godotenv"
)

var app *application

type application struct {
	userSvc    *service.UserService
	contentSvc *service.ContentService
	blRepo     *redisRepository.BlacklistRepository
}

func Initialize() (func(), error) {
	err := godotenv.Load()

	if err != nil {
		return nil, err
	}
	db := postgres.InitDatabase()
	redis := redis.InitRedis()
	blRepo := redisRepository.NewBlacklistRepository(redis)
	userRepo := repository.NewUserRepository(db)
	userSvc := service.NewUserService(userRepo, blRepo)
	contentRepo := repository.NewContentRepository(db)
	contentSvc := service.NewContentService(contentRepo)
	app = &application{
		userSvc:    userSvc,
		contentSvc: contentSvc,
		blRepo:     blRepo,
	}
	cleanup := func() {
		// Close resources here
		fmt.Println("cleaning up ....")
		redis.Close()
		fmt.Println("closed Redis")
	}

	return cleanup, nil

}
