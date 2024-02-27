package protocol

import (
	"os"

	"github.com/brightnc/go-learnhub/internal/adapter/handler/http"
	"github.com/brightnc/go-learnhub/protocol/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ServeREST() {
	userHdl := http.NewUserHandler(app.userSvc)
	contentHdl := http.NewContentHandler(app.contentSvc)
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{os.Getenv("DOMAIN_URL")}
	config.AllowMethods = []string{"GET", "POST", "PATCH", "DELETE"}
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	r.Use(cors.New(config))
	v1 := r.Group("/api/v1")

	v1.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world",
		})
	})
	v1.POST("/register", userHdl.Register)
	v1.POST("/login", userHdl.Login)

	authorized := v1.Group("/")
	authorized.Use(middleware.Auth(app.blRepo))
	authorized.GET("/me", userHdl.GetUserInfo)
	authorized.POST("/logout", userHdl.Logout)
	authorized.POST("/content", contentHdl.CreateContent)
	authorized.PATCH("/content/:id", contentHdl.UpdateContent)
	authorized.DELETE("/content/:id", contentHdl.DeleteContent)

	v1.GET("/content", contentHdl.ContentList)
	v1.GET("/content/:id", contentHdl.GetContentById)

	r.Run(":" + os.Getenv("PORT"))
}
