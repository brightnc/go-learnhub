package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	redisRepository "github.com/brightnc/go-learnhub/internal/adapter/database/redis/repository"
	"github.com/brightnc/go-learnhub/internal/core/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Auth(bl *redisRepository.BlacklistRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")
		if authorizationHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}
		tokenParts := strings.Fields(authorizationHeader)
		if len(tokenParts) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			c.Abort()
			return
		}

		tokenString := tokenParts[1]
		isBlToken := bl.IsInBlackList(tokenString)
		if isBlToken {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		token, err := jwt.ParseWithClaims(tokenString, &service.MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			fmt.Printf("JWT parse error : %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if !token.Valid {
			fmt.Printf("Invalid token : %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		claims, ok := token.Claims.(*service.MyCustomClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unknown claims type"})
			c.Abort()
			return
		}

		c.Set("userId", claims.Id)
		c.Set("expireAt", claims.ExpiresAt.Time)
		c.Set("token", tokenString)

		c.Next()
	}
}
