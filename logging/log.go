package logging

import (
	"face_ui/utils/setting"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"time"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s", setting.ApplicationSetting.LogSavePath)
}

func getLogFileName() string {
	return fmt.Sprintf("%s.%s",
		time.Now().Format(setting.ApplicationSetting.TimeFormat),
		setting.ApplicationSetting.LogFileExt)
}

func Setup() {
	gin.DisableConsoleColor()

	f, _ := os.Create(getLogFilePath() + getLogFileName())
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
