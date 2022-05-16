package repository

import (
	"gorm.io/gorm"
)

type Repo struct {
	Conn *gorm.DB
}
