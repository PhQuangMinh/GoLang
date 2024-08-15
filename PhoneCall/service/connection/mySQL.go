package connection

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQL struct {
	SQL *gorm.DB
}

var mySQL = &MySQL{}

func ConnectDB(user, password, port, host string) *MySQL {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, port, host)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed connect database")
		return nil
	}

	mySQL.SQL = db
	return mySQL
}
