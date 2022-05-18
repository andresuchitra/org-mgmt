package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model

	Text      string `gorm:"type:varchar(255)"`
	IsDeleted bool

	OrganizationID uint `gorm:"index"`
	Organization   Organization
	AuthorID       uint `gorm:"index"`
	Author         User
}

type CreateCommentParam struct {
	OrganizationName string `json:"organization_name"`
	Comment          string `json:"comment"`
	AuthorID         uint   `json:"author_id"`
}
