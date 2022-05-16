package repository

import (
	"context"
	"errors"

	"github.com/andresuchitra/org-mgmt/models"
	"gorm.io/gorm"
)

func NewOrganizationRepository(Conn *gorm.DB) OrganizationRepository {
	return &Repo{Conn}
}

type OrganizationRepository interface {
	GetOrganizationByID(ctx context.Context, id int64) (*models.Organization, error)
}

func (m *Repo) GetOrganizationByID(ctx context.Context, organizationID int64) (*models.Organization, error) {
	if organizationID == 0 {
		return nil, errors.New("Invalid org ID")
	}
	org := models.Organization{}
	result := m.Conn.Where("id = ?", organizationID).First(&org)
	if result.Error != nil {
		return nil, result.Error
	}

	return &org, nil
}
