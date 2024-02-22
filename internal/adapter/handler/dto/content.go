package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateContentRequest struct {
	VideoUrl string `json:"videoUrl" binding:"required"`
	Comment  string `json:"comment" binding:"required"`
	Rating   uint16 `json:"rating" binding:"required"`
}

type ContentResponse struct {
	ID           uuid.UUID `json:"id"`
	VideoTitle   string    `json:"videoTitle"`
	VideoUrl     string    `json:"videoUrl"`
	Comment      string    `json:"comment"`
	Rating       uint16    `json:"rating"`
	ThumbnailUrl string    `json:"thumbnailUrl"`
	CreatorName  string    `json:"creatorName"`
	CreatorUrl   string    `json:"creatorUrl"`
	UserID       uuid.UUID `json:"postedBy"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	DeletedAt    time.Time `json:"deletedAt"`
}
