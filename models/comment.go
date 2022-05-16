package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model

	Text      string `gorm:"type:varchar(255)"`
	IsDeleted bool

	OrganizationID int64 `gorm:"index"`
	Organization   Organization
	AuthorID       int64 `gorm:"index"`
	Author         User
}
