package router

import (
	"face_ui/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": "Register information is not complete",
		})
		return
	}

	isRegistered, err := models.ExistUserByName(user.Username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": "database query error",
		})
		return
	}

	if isRegistered {
		c.JSON(http.StatusOK, gin.H{
			"error": "user exist",
		})
		return
	}

	err = models.AddUser(user)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "database insert error",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	}
}
