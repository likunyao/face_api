package router

import (
	"face_ui/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func tea_register(c *gin.Context) {
	var user models.Teacher
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": "Register information is not complete",
		})
		return
	}

	isRegistered, err := user.ExistUserByName()
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

	err = user.AddUser()
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

func stu_register(c *gin.Context) {
	var user models.Student
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": "Register information is not complete",
		})
		return
	}

	isRegistered, err := user.ExistUserByName()
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

	err = user.AddUser()
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
