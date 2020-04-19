package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Record struct {
	gorm.Model
	Username string `json:"username"`
	Remark   string `json:"remark" gorm:"default:null"`
}

type Result struct {
	CreatedAt   time.Time `json:"date"`
	Username   string    `json:"name"`
	Remark string    `json:"remark"`
}

func GetRecordsFromDatabase(username string) ([]Result, error) {
	var res []Result
	err := db.Table("face_record").Select("created_at, username, remark").Where("username = ?", username).Scan(&res).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return res, nil
}
