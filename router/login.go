package router

import (
	"face_ui/models"
	"face_ui/utils/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error": "Login information is not complete",
		})
		return
	}

	loginSuccess, err := models.AuthorizedByUsernameAndPassword(user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error": "database query error",
		})
		return
	}
	if loginSuccess {
		token, _ := jwt.GenerateToken(user.Username)
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"token": token,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
		})
	}
}
