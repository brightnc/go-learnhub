package repository

import (
	"errors"

	"github.com/brightnc/go-learnhub/internal/core/domain"
	"github.com/google/uuid"
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
	// Fetch the associated user information using Preload
	var resultContent domain.Content
	if err := r.db.Preload("User").First(&resultContent, content.ID).Error; err != nil {
		return nil, err
	}

	// Return the content record with user information
	return &resultContent, nil
}

func (r *ContentRepository) GetContents() ([]*domain.Content, error) {
	var contents []*domain.Content

	result := r.db.Preload("User").Find(&contents)
	if result.Error != nil {
		return nil, result.Error
	}
	return contents, nil

}

func (r *ContentRepository) GetContentById(id string) (*domain.Content, error) {
	var content *domain.Content

	result := r.db.Preload("User").First(&content, "id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("content not found")
	}

	return content, nil
}

func (r *ContentRepository) UpdateContent(id string, content *domain.Content) (*domain.Content, error) {
	result := r.db.Where("id=?", id).Updates(content)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("cannot update content")
	}

	var resultContent domain.Content
	contentUuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	if err := r.db.Preload("User").First(&resultContent, contentUuid).Error; err != nil {
		return nil, err
	}

	return &resultContent, nil
}

func (r *ContentRepository) DeleteContent(id string) error {
	contentUuid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	result := r.db.Delete(&domain.Content{}, contentUuid)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("cannot delete content")
	}
	return nil
}
