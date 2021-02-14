package models

import (
	"zhouyongProject/conf"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var db *gorm.DB

func init() {
	var err error
	db,err = gorm.Open(conf.DBType,conf.DbUrl)
	if err != nil {
		log.Println(err)
	}
	db.SingularTable(true)
	db.AutoMigrate(&StudentAuth{},&SelectSheet{},&SelectChar{},&SelectWord{},&SelectStock{},
		&HomeworkDoc{},&StudentHomework{},&Student{},&StudentHomeworkResult{},
		&WordError{},&Teacher{})
}
