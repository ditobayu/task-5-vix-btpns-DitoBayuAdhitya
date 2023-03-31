package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/task_5_vix_btpns"))
	if err != nil {
		panic(err)
	}
	database.AutoMigrate(&User{})
	database.AutoMigrate(&Photo{})

	DB = database
}