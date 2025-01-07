package db

import (
	"yusrilaprial/backend-api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase() {
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/db_go_api"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Post{})
	DB = db
}
