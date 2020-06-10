package models

import (
	"face_ui/utils/setting"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var db *gorm.DB

//Setup initialize database face_api and tables: user, record
func Setup() {
	var err error
	db, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name))
	if err != nil {
		log.Fatalf("models.Setup failed, err: %v", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DatabaseSetting.TablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.LogMode(true)

	if db.HasTable(&Student{}) {
		db.AutoMigrate(&Student{})
	} else {
		db.CreateTable(&Student{})
	}

	if db.HasTable(&Teacher{}) {
		db.AutoMigrate(&Teacher{})
	} else {
		db.CreateTable(&Teacher{})
	}

	if db.HasTable(&Record{}) {
		db.AutoMigrate(&Record{})
	} else {
		db.CreateTable(&Record{})
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}
func CloseDB() {
	defer db.Close()
}
