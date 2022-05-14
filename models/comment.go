package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model

	Text      string `gorm:"type:varchar(255)"`
	IsDeleted bool

	OrganizationID int64
	Organization   Organization
	MemberID       int64
	Author         User
}
