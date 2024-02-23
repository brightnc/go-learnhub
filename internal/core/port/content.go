package port

import "github.com/brightnc/go-learnhub/internal/core/domain"

type IContentRepository interface {
	CreateContent(content *domain.Content) (*domain.Content, error)
	GetContents() ([]*domain.Content, error)
	GetContentById(id string) (*domain.Content, error)
	UpdateContent(id string, content *domain.Content) (*domain.Content, error)
	DeleteContent(id string) error
}

type IContentService interface {
	CreateContent(userId string, videoUrl string, comment string, rating uint16) (*domain.Content, error)
	GetContents() ([]*domain.Content, error)
	GetContentById(id string) (*domain.Content, error)
	UpdateContent(id string, content *domain.Content) (*domain.Content, error)
	DeleteContent(id string) error
}
