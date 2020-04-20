package controller

import (
	"face_ui/models"
	"github.com/gin-gonic/gin"
	jsontime "github.com/liamylian/jsontime/v2/v2"
	"net/http"
	"time"
)

var json = jsontime.ConfigWithCustomTimeFormat

func init() {
	jsontime.SetDefaultTimeFormat("2006-01-02 15:04:05", time.Local)
}

func GetRecords(c *gin.Context) {
	data, ok := c.Get("username")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"err": "token invalid",
		})
		return
	}
	username, _ := data.(string)
	records, err := models.GetRecordsFromDatabase(username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"err": "database query err",
		})
		return
	}
	jsonRecords, _ := json.Marshal(records)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": string(jsonRecords),
	})
}
