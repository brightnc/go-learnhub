package http

import (
	"net/http"

	"github.com/brightnc/go-learnhub/internal/adapter/handler/dto"
	"github.com/brightnc/go-learnhub/internal/core/port"
	"github.com/gin-gonic/gin"
)

type ContentHandler struct {
	svc port.IContentService
}

func NewContentHandler(svc port.IContentService) *ContentHandler {
	return &ContentHandler{
		svc: svc,
	}
}

func (contentHdl *ContentHandler) CreateContent(c *gin.Context) {
	var req dto.CreateContentRequest
	userId := c.MustGet("userId").(string)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}

	result, err := contentHdl.svc.CreateContent(userId, req.VideoUrl, req.Comment, uint16(req.Rating))
	if err != nil {
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "cannot create content",
			})
		}
		return
	}

	response := dto.ContentResponse{
		ID:           result.ID,
		VideoTitle:   result.VideoTitle,
		VideoUrl:     result.VideoUrl,
		Comment:      result.Comment,
		Rating:       result.Rating,
		ThumbnailUrl: result.ThumbnailUrl,
		CreatorName:  result.CreatorName,
		CreatorUrl:   result.CreatorUrl,
		UserID:       result.UserID,
		CreatedAt:    result.CreatedAt,
		UpdatedAt:    result.UpdatedAt,
		DeletedAt:    result.DeletedAt.Time,
	}
	c.JSON(http.StatusOK, response)
}
