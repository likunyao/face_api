package models

import (
	"crypto/md5"
	"fmt"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username  string `json:"username" gorm:"unique; not null"`
	Password  string `json:"password"`
	Faceimage string `json:"faceimage"`
	Role      int    `json:"role" gorm:"default:0"`
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

func AddUser(data User) error {
	p := md5.Sum([]byte(data.Password))
	user := User{
		Username:  data.Username,
		Password:  fmt.Sprintf("%x", p),
		Faceimage: data.Faceimage,
		Role:      0,
	}
	err := db.Create(&user).Error

	return err
}

func AuthorizedByUsernameAndPassword(data User) (bool, error) {
	var user User
	err := db.Where("username = ?", data.Username).First(&user).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	p := md5.Sum([]byte(data.Password))
	if fmt.Sprintf("%x", p) == user.Password {
		return true, nil
	}
	return false, nil
}
