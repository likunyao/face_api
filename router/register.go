package router

import (
	"face_ui/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Login information is not complete",
		})
		return
	}
	err := models.AddUser(user)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "user exist",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	}
}
