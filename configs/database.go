package configs

import (
	"go-login-api/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=postgres dbname=go_login_api port=5432 sslmode=disable TimeZone=Asia/Jakarta"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.MUser{})

	DB = db
	log.Println("Database connected")
}
