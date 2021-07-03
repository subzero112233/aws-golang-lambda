package middleware

import (
	"fmt"

	"aws-golang-lambda/pkg/jwtparser"

	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		const prefix = "Bearer "

		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatusJSON(400, gin.H{"error": "Authorization header missing or empty"})
		}

		token := header[len(prefix):]
		claims, err := jwtparser.ValidateToken(token, secret)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": fmt.Sprintf("unable to verify token with error: %v", err)})
		}

		c.Set("username", claims.Username)

		c.Next()
	}
}
