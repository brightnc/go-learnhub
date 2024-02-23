package dto

import (
	"time"

	"github.com/brightnc/go-learnhub/internal/core/domain"
	"github.com/google/uuid"
)

type CreateContentRequest struct {
	VideoUrl string `json:"videoUrl" binding:"required"`
	Comment  string `json:"comment" binding:"required"`
	Rating   uint16 `json:"rating" binding:"required"`
}

type ContentResponse struct {
	ID           uuid.UUID    `json:"id"`
	VideoTitle   string       `json:"videoTitle"`
	VideoUrl     string       `json:"videoUrl"`
	Comment      string       `json:"comment"`
	Rating       uint16       `json:"rating"`
	ThumbnailUrl string       `json:"thumbnailUrl"`
	CreatorName  string       `json:"creatorName"`
	CreatorUrl   string       `json:"creatorUrl"`
	User         UserResponse `json:"postedBy"`
	CreatedAt    time.Time    `json:"createdAt"`
	UpdatedAt    time.Time    `json:"updatedAt"`
	DeletedAt    time.Time    `json:"deletedAt"`
}

func ToContentResponse(content *domain.Content) ContentResponse {
	return ContentResponse{
		ID:           content.ID,
		VideoTitle:   content.VideoTitle,
		VideoUrl:     content.VideoUrl,
		Comment:      content.Comment,
		Rating:       content.Rating,
		ThumbnailUrl: content.ThumbnailUrl,
		CreatorName:  content.CreatorName,
		CreatorUrl:   content.CreatorUrl,
		User:         ToUserResponse(content.User),
		CreatedAt:    content.CreatedAt,
		UpdatedAt:    content.UpdatedAt,
		DeletedAt:    content.DeletedAt.Time,
	}
}
