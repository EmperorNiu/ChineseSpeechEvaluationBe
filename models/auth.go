package models

import (
	_ "github.com/jinzhu/gorm"
)

// 学生账户管理
type StudentAuth struct {
	AuthId uint `json:"auth_id" gorm:"primary_key;auto-increment"`
	Username string `json:"username" form:"username" gorm:"uniqueIndex"`
	Password string `json:"password" form:"password"`
	StudentIdRefer string `json:"student_id_refer" form:"student_id_refer"`
	Student Student `gorm:"foreignkey:StudentId; references:StudentIdRefer"`
}


func (user *StudentAuth) Create() error{
	stu := Student{StudentId: user.StudentIdRefer, Name: user.Username}
	err := db.Create(&stu).Error
	if err != nil {
		return err
	} else {
		err2 := db.Create(&user).Error
		return err2
	}
}

func (user *StudentAuth) Query(name string) error {
	err := db.Where(&StudentAuth{Username: name}).First(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (user *StudentAuth) Update(attribute,value string) {
	db.Model(user).Update(attribute,value)
}
