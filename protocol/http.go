package protocol

import (
	"github.com/brightnc/go-learnhub/internal/adapter/handler/http"
	"github.com/brightnc/go-learnhub/protocol/middleware"
	"github.com/gin-gonic/gin"
)

func ServeREST() {
	userHdl := http.NewUserHandler(app.userSvc)
	contentHdl := http.NewContentHandler(app.contentSvc)
	r := gin.Default()
	v1 := r.Group("/api/v1")

	v1.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world",
		})
	})
	v1.POST("/register", userHdl.Register)
	v1.POST("/login", userHdl.Login)

	authorized := v1.Group("/auth")
	authorized.Use(middleware.Auth())
	authorized.GET("/mine", userHdl.GetUserById)
	authorized.POST("/content", contentHdl.CreateContent)

	v1.GET("/content", contentHdl.ContentList)
	v1.GET("/content/:id", contentHdl.GetContentById)

	r.Run(":8000")
}
