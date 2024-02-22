package protocol

import (
	"log"

	"github.com/brightnc/go-learnhub/internal/adapter/database/postgres"
	"github.com/brightnc/go-learnhub/internal/adapter/database/postgres/repository"
	"github.com/brightnc/go-learnhub/internal/core/service"
	"github.com/joho/godotenv"
)

var app *application

type application struct {
	userSvc    *service.UserService
	contentSvc *service.ContentService
}

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading env")
	}
	db := postgres.InitDatabase()
	userRepo := repository.NewUserRepository(db)
	userSvc := service.NewUserService(userRepo)
	contentRepo := repository.NewContentRepository(db)
	contentSvc := service.NewContentService(contentRepo)
	app = &application{
		userSvc:    userSvc,
		contentSvc: contentSvc,
	}

}
