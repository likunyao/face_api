package router

import (
	"face_ui/models"
	"face_ui/utils/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func tea_login(c *gin.Context) {
	var user models.Teacher
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   "Login information is not complete",
		})
		return
	}

	loginSuccess, err := user.AuthorizedByUsernameAndPassword()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   "database query error",
		})
		return
	}
	if loginSuccess {
		token, _ := jwt.GenerateToken(user.Username)
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"token":   token,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
		})
	}
}

func stu_login(c *gin.Context) {
	var user models.Student
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   "Login information is not complete",
		})
		return
	}

	loginSuccess, err := user.AuthorizedByUsernameAndPassword()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   "database query error",
		})
		return
	}
	if loginSuccess {
		token, _ := jwt.GenerateToken(user.Username)
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"token":   token,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
		})
	}
}
