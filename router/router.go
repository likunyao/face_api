package router

import (
	"face_ui/middleware"
	"face_ui/router/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	r.POST("/register",register)
	r.POST("/login", login)
	r.GET("/records", middleware.Auth(), controller.GetRecords)

	return r
}
