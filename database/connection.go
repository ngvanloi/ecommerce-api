package database

import (
	"log"

	"github.com/ngvanloi/ecommerce-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("ecommerce.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}

	// Tự động tạo bảng nếu chưa tồn tại
	DB.AutoMigrate(&models.Product{})
}
