package models

import (
	"crypto/md5"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username  string `json:"username" gorm:"unique; not null"`
	Password  string `json:"password"`
	Faceimage string `json:"faceimage"`
	Role      int    `gorm:"default:0"`
}

func ExistUserByName(name string) (bool, error) {
	var user User
	err := db.Where("username = ?", name).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if user.ID > 0 {
		return true, nil
	}

	return false, nil
}

func AddUser(data map[string]interface{}) error {
	isRegistered, err := ExistUserByName(data["username"].(string))
	if err != nil {
		panic("database error")
	}
	if isRegistered {

	}

	p := md5.Sum([]byte(data["password"].(string)))
	user := User{
		Username:  data["username"].(string),
		Password:  string(p[:]),
		Faceimage: "",
		Role:      0,
	}
	err = db.Create(&user).Error

	return err
}
