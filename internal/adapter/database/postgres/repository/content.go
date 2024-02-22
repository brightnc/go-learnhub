package repository

import (
	"errors"

	"github.com/brightnc/go-learnhub/internal/core/domain"
	"gorm.io/gorm"
)

type ContentRepository struct {
	db *gorm.DB
}

func NewContentRepository(db *gorm.DB) *ContentRepository {
	return &ContentRepository{
		db: db,
	}
}

func (r *ContentRepository) CreateContent(content *domain.Content) (*domain.Content, error) {
	result := r.db.Create(content)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("cannot create content")
	}

	return &domain.Content{
		ID:           content.ID,
		VideoTitle:   content.VideoTitle,
		VideoUrl:     content.VideoUrl,
		Comment:      content.Comment,
		Rating:       content.Rating,
		ThumbnailUrl: content.ThumbnailUrl,
		CreatorName:  content.CreatorName,
		CreatorUrl:   content.CreatorUrl,
		UserID:       content.UserID,
		CreatedAt:    content.CreatedAt,
		UpdatedAt:    content.UpdatedAt,
		DeletedAt:    content.DeletedAt,
	}, nil
}

func (r *ContentRepository) GetContents() ([]*domain.Content, error) {
	var contents []*domain.Content

	result := r.db.Find(&contents)
	if result.Error != nil {
		return nil, result.Error
	}
	return contents, nil

}

func (r *ContentRepository) GetContentById(id string) (*domain.Content, error) {
	var content *domain.Content

	result := r.db.First(&content, "id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("content not found")
	}

	return content, nil
}
