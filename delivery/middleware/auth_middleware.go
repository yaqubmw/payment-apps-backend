package middleware

import (
	"fmt"
	"net/http"
	"payment-apps-backend/security"
	"strings"

	"github.com/gin-gonic/gin"
)

type authHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

func AuthMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var h authHeader
		if err := c.ShouldBindHeader(&h); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			c.Abort()
			return
		}

		tokenHeader := strings.Replace(h.AuthorizationHeader, "Bearer ", "", -1)
		fmt.Println("tokenHeader:", tokenHeader)

		if tokenHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			c.Abort()
			return
		}

		token, err := security.VerifyAccessToken(tokenHeader)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			c.Abort()
			return
		}
		if token != nil {
			if role, ok := token["role"].(string); !ok || role != requiredRole {
				c.JSON(http.StatusForbidden, gin.H{"message": "forbidden"})
				c.Abort()
				return
			}

			c.Set("token", token)
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			c.Abort()
			return
		}
	}
}
