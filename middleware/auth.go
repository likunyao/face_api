package middleware

import (
	"face_ui/utils/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if len(token) == 0 {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"err": "token not found",
			})
			c.Abort()
		}

		username, err := jwt.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"err": "token invalid",
			})
			c.Abort()
		}
		c.Set("username", username)
	}
}
