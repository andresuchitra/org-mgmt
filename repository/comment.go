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

func (m *Repo) GetCommentsByOrgID(ctx context.Context, organizationID int64) ([]models.Comment, error) {
	comments := make([]models.Comment, 0)

	result := m.Conn.Where("organization_id = ?", organizationID).Find(&comments)
	if result.Error != nil {
		return nil, result.Error
	}

	return comments, nil
}
