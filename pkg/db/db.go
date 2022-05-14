package db

import (
	"fmt"
	"log"
	"os"

	"github.com/andresuchitra/org-mgmt/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	// setuo DB connection
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Organization{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Comment{})
	db.AutoMigrate(&models.Follower{})

	return db
}
