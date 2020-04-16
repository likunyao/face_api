package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Record struct {
	gorm.Model
	Username string
	Time time.Time
	Remark string
}