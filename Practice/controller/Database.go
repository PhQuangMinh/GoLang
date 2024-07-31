package controller

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MakeGorm() *gorm.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/call_management?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed connect database")
		return nil
	}
	return db
}
