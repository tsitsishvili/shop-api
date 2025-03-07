package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tsitsishvili/shop-api/internal/repository"
)

func APIKeyMiddleware(shopRepo *repository.ShopRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing API Key"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, "ApiKey ")
		if len(parts) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API Key Format"})
			c.Abort()
			return
		}

		apiKey := parts[1]
		shop, err := shopRepo.FindByAPIKey(apiKey)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API Key"})
			c.Abort()
			return
		}

		c.Set("shopID", shop.ID)
		c.Next()
	}
}
