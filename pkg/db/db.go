package db

import (
	"errors"
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

	log.Println("Connecting to: ", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	// seed organization first record
	firstOrg := models.Organization{Name: "Xendit"}
	if err = db.AutoMigrate(&models.Organization{}); err == nil && db.Migrator().HasTable(&models.Organization{}) {
		if err := db.First(&models.Organization{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			firstOrg = models.Organization{Name: "xendit"}
			db.Create(&firstOrg)
		} else if err != nil {
			log.Fatalln("Error seeding organization!")
		}
	}
	// seed user first record
	if err = db.AutoMigrate(&models.User{}); err == nil && db.Migrator().HasTable(&models.User{}) {
		if err := db.First(&models.User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			data := models.User{
				Name:           "First",
				Email:          "em@ail.com",
				Password:       "pass123",
				OrganizationID: 1, // hardcode to 1, assuming organization table start with ID 1
			}
			db.Create(&data)
		}
	}

	db.AutoMigrate(&models.Comment{})

	return db
}
