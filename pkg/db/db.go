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

	log.Println("Connecting to DB: ", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	// seed organization first record
	if err := db.AutoMigrate(&models.Organization{}); err == nil && db.Migrator().HasTable(&models.Organization{}) {
		if err := db.First(&models.Organization{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			orgs := []models.Organization{
				{Name: "xendit"},
				{Name: "midtrans"},
			}
			db.CreateInBatches(orgs, 2)
		}
	}

	// seed users
	if err := db.AutoMigrate(&models.User{}); err == nil && db.Migrator().HasTable(&models.User{}) {
		checkErr := db.First(&models.User{}).Error

		if checkErr != nil {
			if errors.Is(checkErr, gorm.ErrRecordNotFound) {
				users := []models.User{
					{
						Name:           "First User",
						Email:          "em@ail.com",
						Password:       "pass123",
						OrganizationID: 1,
					},
					{
						Name:           "Second User",
						Email:          "asd@qw.com",
						Password:       "qqq123",
						OrganizationID: 2,
					},
					{
						Name:           "THIRD User",
						Email:          "iii@ail.com",
						Password:       "pas222",
						OrganizationID: 1,
					},
				}
				db.CreateInBatches(users, 3)
			} else {
				log.Fatalln(checkErr.Error())
			}
		}
	}

	// seed user_followers
	db.Exec("INSERT INTO user_followers(user_id, follower_id) VALUES(1,2)")
	db.Exec("INSERT INTO user_followers(user_id, follower_id) VALUES(1,3)")
	db.Exec("INSERT INTO user_followers(user_id, follower_id) VALUES(2,1)")
	db.Exec("INSERT INTO user_followers(user_id, follower_id) VALUES(2,3)")
	db.Exec("INSERT INTO user_followers(user_id, follower_id) VALUES(3,2)")

	// seed comments
	if err := db.AutoMigrate(&models.Comment{}); err == nil && db.Migrator().HasTable(&models.Comment{}) {
		if err := db.First(&models.Comment{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			data := []models.Comment{
				{Text: "Comment on xendit 1", OrganizationID: 1, AuthorID: 1},
				{Text: "Comment by Third User", OrganizationID: 1, AuthorID: 3},
				{Text: "This is midtrans", OrganizationID: 2, AuthorID: 2},
			}
			db.CreateInBatches(data, 3)
		}
	}

	return db
}
