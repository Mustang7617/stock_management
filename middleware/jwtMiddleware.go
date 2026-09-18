package middleware

import (
	"api/config"
	"api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// var jwtKey = []byte("secret_key") // Fix: "sercret_key" corrected to "secret_key"

// func AuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		authHeader := c.GetHeader("Authorization")
// 		if authHeader == "" {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
// 			c.Abort()
// 			return
// 		}

// 		parts := strings.Split(authHeader, "Bearer ")
// 		if len(parts) != 2 {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
// 			c.Abort()
// 			return
// 		}

// 		tokenString := parts[1]

// 		token, _ := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
// 			return jwtKey, nil
// 		})

// 		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 			c.Set("user", claims) // Example: setting claims in the context for later use
// 			c.Next()
// 		} else {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
// 			c.Abort()
// 		}
// 	}
// }

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization") // Example token-based auth
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			c.Abort()
			return
		}

		// Assuming token contains the user’s ID or some other identifier
		var user model.User
		if err := config.DB.Where("token = ?", token).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Store the authenticated user in the context
		c.Set("user", user)
		c.Next()
	}
}
