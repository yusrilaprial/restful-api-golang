package db

import (
	"yusrilaprial/backend-api/configs"
	"yusrilaprial/backend-api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := configs.DB_USERNAME + ":" + configs.DB_PASSWORD + "@tcp(" + configs.DB_HOST + ":" + configs.DB_PORT + ")/" + configs.DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Post{})
	DB = db
}
