package data

import (
	"bookstore/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error
	dbname := "root:admin@tcp/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dbname), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if err = DB.AutoMigrate(&model.User{}, &model.Book{}); err != nil {
		panic(err)
	}

}
