package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Name     string `gorm:"type:varchar(255)"`
	Email    string `gorm:"type:varchar(255);not null"`
	Password string `gorm:"type:varchar(16);not null"`
	Avatar   string `gorm:"type:varchar(255)"`

	Followers      []User `gorm:"many2many:user_followers"`
	OrganizationID int64  `gorm:"index;not null"`
	Organization   Organization
}

type UserWithFollower struct {
	User

	TotalFollowing uint32 `gorm:"->"`
	TotalFollower  uint32 `gorm:"->"`
}
