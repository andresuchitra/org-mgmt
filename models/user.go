package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Name     string `gorm:"type:varchar(255)"`
	Email    string `gorm:"type:varchar(255);not null"`
	Password string `gorm:"type:varchar(16);not null"`
	Avatar   string `gorm:"type:varchar(255)"`
}

type Follower struct {
	gorm.Model

	UserID     int64 `gorm:"index;not null"`
	FollowerID int64 `gorm:"index;not null"`
}
