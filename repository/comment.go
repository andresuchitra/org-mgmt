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
	GetCommentsByOrgID(ctx context.Context, id uint) ([]models.Comment, error)
	CreateCommentByOrganizationID(ctx context.Context, organizationID uint, authorID uint, comment string) error
	SoftDeleteCommentsByOrganizationID(ctx context.Context, organizationID uint) error
}

func (m *Repo) GetCommentsByOrgID(ctx context.Context, organizationID uint) ([]models.Comment, error) {
	comments := make([]models.Comment, 0)

	result := m.Conn.Where("organization_id = ? AND is_deleted = false", organizationID).Find(&comments)
	if result.Error != nil {
		return nil, result.Error
	}

	return comments, nil
}

func (m *Repo) CreateCommentByOrganizationID(ctx context.Context, organizationID uint, authorID uint, comment string) error {
	newComment := models.Comment{
		Text:           comment,
		OrganizationID: organizationID,
		AuthorID:       authorID,
	}

	result := m.Conn.Create(&newComment)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *Repo) SoftDeleteCommentsByOrganizationID(ctx context.Context, organizationID uint) error {
	result := m.Conn.Table("comments").Where("organization_id = ?", organizationID).Update("is_deleted", true)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
