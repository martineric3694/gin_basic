package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/golang"))

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Department{})

	DB = db
}
