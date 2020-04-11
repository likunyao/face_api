package main

import (
	"face_ui/logging"
	"face_ui/models"
	"face_ui/router"
	"face_ui/utils/setting"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func init() {
	setting.Setup()
	models.Setup()
	logging.Setup()
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routerInit := router.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeOut
	writeTiemout := setting.ServerSetting.WriteTimeOut
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:              endPoint,
		Handler:           routerInit,
		ReadTimeout:       readTimeout,
		WriteTimeout:      writeTiemout,
		MaxHeaderBytes:    maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)
	server.ListenAndServe()
}