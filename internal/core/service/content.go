package service

import (
	"github.com/brightnc/go-learnhub/internal/core/domain"
	"github.com/brightnc/go-learnhub/internal/core/port"
	"github.com/brightnc/go-learnhub/internal/core/util"
	"github.com/google/uuid"
)

type ContentService struct {
	repo port.IContentRepository
}

func NewContentService(repo port.IContentRepository) *ContentService {
	return &ContentService{
		repo: repo,
	}
}

func (svc *ContentService) CreateContent(userId string, videoUrl string, comment string, rating uint16) (*domain.Content, error) {
	oembedData, err := util.Oembed(videoUrl)
	if err != nil {
		return nil, err
	}
	userUuid, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}
	newContent := &domain.Content{
		VideoUrl:     videoUrl,
		Comment:      comment,
		Rating:       rating,
		CreatorName:  oembedData.Author_name,
		CreatorUrl:   oembedData.Author_url,
		UserID:       userUuid,
		ThumbnailUrl: oembedData.Thumbnail_url,
		VideoTitle:   oembedData.Title,
	}
	result, err := svc.repo.CreateContent(newContent)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (svc *ContentService) GetContents() ([]*domain.Content, error) {
	return nil, nil
}
func (svc *ContentService) GetContentById(id string) (*domain.Content, error) {
	return nil, nil
}
