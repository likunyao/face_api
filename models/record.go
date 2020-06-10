package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Record struct {
	gorm.Model
	Username  string `json:"username"`
	StudentID int    `json:"studentid" gorm:"unique; not null"`
	Remark    string `json:"remark" gorm:"default:null"`
}

type RecordResult struct {
	CreatedAt time.Time `json:"datetime"`
	Username  string    `json:"name"`
	StudentID int       `json:"studentid"`
	Remark    string    `json:"remark"`
}

func GetRecordsFromDatabase(username string) ([]RecordResult, error) {
	var res []RecordResult
	err := db.Table("face_record").Select("created_at, username, remark").Where("username = ?", username).Scan(&res).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return res, nil
}
