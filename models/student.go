package models

import (
	"crypto/md5"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Student struct {
	gorm.Model
	Username  string `json:"username" gorm:"unique; not null"`
	Password  string `json:"password"`
	StudentID int    `json:"studentid" gorm:"unique; not null"`
	TeacherID int    `json:"teacherid" gorm:"unique; not null"`
	Imgpath   string `json:"imgpath"`
}

//type User struct {
//	gorm.Model
//	Username  string `json:"username" gorm:"unique; not null"`
//	Password  string `json:"password"`
//	Faceimage string `json:"faceimage" gorm:"not null"`
//	Role      int    `json:"role" gorm:"default:0"`
//}

func (stu *Student) ExistUserByName() (bool, error) {
	var user Student
	err := db.Where("username = ?", stu.Username).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if user.ID > 0 {
		return true, nil
	}

	return false, nil
}

func (stu *Student) AddUser() error {
	p := md5.Sum([]byte(stu.Password))
	user := Student{
		Username:  stu.Username,
		Password:  fmt.Sprintf("%x", p),
		StudentID: stu.StudentID,
		TeacherID: stu.TeacherID,
		Imgpath:   stu.Imgpath,
	}
	err := db.Create(&user).Error

	return err
}

func (stu *Student) AuthorizedByUsernameAndPassword() (bool, error) {
	var user Student
	err := db.Where("username = ?", stu.Username).First(&user).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	p := md5.Sum([]byte(stu.Password))
	if fmt.Sprintf("%x", p) == user.Password {
		return true, nil
	}
	return false, nil
}

func (stu *Student) GetRecords() ([]RecordResult, error) {
	var res []RecordResult
	err := db.Table("face_record").Select("created_at, username, remark").Where("username = ?", stu.Username).Scan(&res).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return res, nil
}
