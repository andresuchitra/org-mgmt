package models

import "gorm.io/gorm"

type Organization struct {
	gorm.Model

	Name    string `gorm:"type:varchar(255);uniqueIndex;not null"`
	Members []*User
}
