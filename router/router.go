package router

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/register",register)
	r.POST("/login", login)

	return r
}
