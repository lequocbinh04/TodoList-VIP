package middleware

import (
	"TodoList/JWT"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func (c *gin.Context) {
		clientToken := c.Request.Header.Get("Authorization")
		if clientToken == "" {
			c.JSON(403, gin.H{
				"message": "No Authorization header provided",
				"status": "false",
			})
			c.Abort()
			return
		}

		extractedToken := strings.Split(clientToken, "Bearer ")
		if len(extractedToken) == 2 {
			clientToken = strings.TrimSpace(extractedToken[1])
		} else {
			c.JSON(400, gin.H{
				"message": "Incorrect Format of Authorization Token",
				"status": "false",
			})
			c.Abort()
			return
		}
		jwtWrapper := JWT.JwtWrapper{
			SecretKey:       "IGhr2aIe6qnu8qBqXkC5X6DFbI4xxXQ7",
			Issuer:          "AuthService",
			ExpirationHours: 1000000000000000,
		}
		
		claims, err := jwtWrapper.ValidateToken(clientToken)
		if err != nil {
			c.JSON(401, gin.H{
				"message": "Unauthorized user",
				"status": "false",
				"error": err.Error(),
			})
			c.Abort()
			return
		}
		c.Set("logged_address", claims.Address)
		c.Next()
	}
}
