package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/brightnc/go-learnhub/internal/adapter/handler/dto"
	"github.com/brightnc/go-learnhub/internal/core/domain"
	"github.com/brightnc/go-learnhub/internal/core/port"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc port.IUserService
}

func NewUserHandler(svc port.IUserService) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

func (userHdl *UserHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}
	user := domain.User{
		Name:     req.Name,
		Username: req.Username,
		Password: req.Password,
	}

	result, err := userHdl.svc.Register(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "cannot register",
		})
		return
	}
	fmt.Println(result)
	response := dto.RegisterResponse{
		ID:        result.ID,
		Name:      result.Name,
		Username:  result.Username,
		CreatedAt: result.CreatedAt,
	}
	c.JSON(http.StatusCreated, response)
}

func (userHdl *UserHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}
	token, err := userHdl.svc.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "cannot login",
		})
		return
	}
	response := dto.LoginResponse{
		Token: token,
	}
	c.JSON(http.StatusOK, response)
}

func (userHdl *UserHandler) Logout(c *gin.Context) {
	expireAt := c.MustGet("expireAt").(time.Time)
	token := c.MustGet("token").(string)
	err := userHdl.svc.Logout(token, expireAt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "cannot logout",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "logout success",
	})
}

func (userHdl *UserHandler) GetUserInfo(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	result, err := userHdl.svc.GetUserById(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "wrong id",
		})
	}
	fmt.Println(result)
	response := dto.UserResponse{
		ID:        result.ID,
		Name:      result.Name,
		Username:  result.Username,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}
	c.JSON(http.StatusOK, response)
}

func (userHdl *UserHandler) TestAuth(c *gin.Context) {
	userId := c.MustGet("userId")
	fmt.Printf("user-id : %v", userId)
	c.JSON(http.StatusOK, userId)
}
