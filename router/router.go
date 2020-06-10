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

	r.POST("/sturegister", stu_register)
	r.POST("/tearegister", tea_register)
	r.POST("/stulogin", stu_login)
	r.POST("/tealogin", tea_login)
	r.GET("/records", middleware.Auth(), controller.GetRecords)

	return r
}
