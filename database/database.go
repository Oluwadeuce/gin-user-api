package database

import (
	"gin-api-portfolio/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Database() {

	database, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	database.AutoMigrate(&models.User{})

	DB = database
}
