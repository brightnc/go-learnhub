package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Content struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	VideoTitle   string
	VideoUrl     string
	Comment      string
	Rating       uint16
	ThumbnailUrl string
	CreatorName  string
	CreatorUrl   string
	UserID       uuid.UUID
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
