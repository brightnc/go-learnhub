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

	response := dto.ToContentResponse(result)
	c.JSON(http.StatusOK, response)
}

func (contentHdl *ContentHandler) ContentList(c *gin.Context) {
	result, err := contentHdl.svc.GetContents()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "cannot list content",
		})
		return
	}

	response := make([]dto.ContentResponse, len(result))
	for i, content := range result {
		response[i] = dto.ToContentResponse(content)
	}

	c.JSON(http.StatusOK, response)

}

func (contentHdl *ContentHandler) GetContentById(c *gin.Context) {
	id := c.Param("id")

	result, err := contentHdl.svc.GetContentById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "wrong content id",
		})
		return
	}
	response := dto.ToContentResponse(result)

	c.JSON(http.StatusOK, response)
}
