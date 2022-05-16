package repository

import (
	"context"

	"github.com/andresuchitra/org-mgmt/models"
	"gorm.io/gorm"
)

func NewCommentRepository(Conn *gorm.DB) CommentRepository {
	return &Repo{Conn}
}

type CommentRepository interface {
	GetCommentsByOrgID(ctx context.Context, id int64) ([]models.Comment, error)
}

func (m *Repo) GetCommentsByOrgID(ctx context.Context, organizationID int64) ([]models.Comment, err error) {
	
}
