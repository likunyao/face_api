package models

import (
	"crypto/md5"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Teacher struct {
	gorm.Model
	Username  string `json:"username" gorm:"unique; not null"`
	Password  string `json:"password"`
	College   string `json:"college" gorm:"not null"`
	TeacherID int    `json:"teacherid" gorm:"unique; not null"`
}

func (tea *Teacher) ExistUserByName() (bool, error) {
	var user Teacher
	err := db.Where("username = ?", tea.Username).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if user.ID > 0 {
		return true, nil
	}

	return false, nil
}

func (tea *Teacher) AddUser() error {
	p := md5.Sum([]byte(tea.Password))
	user := Teacher{
		Username:  tea.Username,
		Password:  fmt.Sprintf("%x", p),
		College:   tea.College,
		TeacherID: tea.TeacherID,
	}
	err := db.Create(&user).Error

	return err
}

func (tea *Teacher) AuthorizedByUsernameAndPassword() (bool, error) {
	var user Teacher
	err := db.Where("username = ?", tea.Username).First(&user).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	p := md5.Sum([]byte(tea.Password))
	if fmt.Sprintf("%x", p) == user.Password {
		return true, nil
	}
	return false, nil
}

func (tea *Teacher) GetRecords() ([]RecordResult, error) {
	var res []RecordResult
	return res, nil
}
