package controller

import (
	"encoding/json"
	"face_ui/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
