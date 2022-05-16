package repository

import (
	"context"

	"github.com/andresuchitra/org-mgmt/models"
	"gorm.io/gorm"
)

type repo struct {
	Conn *gorm.DB
}

type CommentRepository interface {
	GetCommentsByOrgID(ctx context.Context, id int64) ([]models.Comment, error)
}
type OrganizationRepository interface {
	GetByID(ctx context.Context, id int64) (models.Organization, error)
}
type UserRepository interface {
	GetMembersByOrganizationID(ctx context.Context, organizationID int64) ([]models.User, error)
}

func NewRepository(Conn *gorm.DB) CommentRepository {
	return &repo{Conn}
}

func (m *repo) GetCommentsByOrgID(ctx context.Context, organizationID) ([]models.Comment, err error) {

}
