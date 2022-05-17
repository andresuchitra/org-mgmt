package repository

import (
	"context"

	"github.com/andresuchitra/org-mgmt/models"
	"gorm.io/gorm"
)

func NewUserRepository(Conn *gorm.DB) UserRepository {
	return &Repo{Conn}
}

type UserRepository interface {
	GetMembersByOrganizationID(ctx context.Context, organizationID int64) ([]models.User, error)
}

func (m *Repo) GetMembersByOrganizationID(ctx context.Context, organizationID int64) ([]models.User, error) {
	members := make([]models.User, 0)

	result := m.Conn.Where("organization_id = ?", organizationID).Find(members)
	if result.Error != nil {
		return nil, result.Error
	}

	return members, nil
}
