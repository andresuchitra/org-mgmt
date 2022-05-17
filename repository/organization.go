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
	GetOrganizationByName(ctx context.Context, name string) (*models.Organization, error)
}

func (m *Repo) GetOrganizationByName(ctx context.Context, name string) (*models.Organization, error) {
	if name == "" {
		return nil, errors.New("Invalid org name")
	}
	org := models.Organization{}
	result := m.Conn.Where("name = ?", name).First(&org)
	if result.Error != nil {
		return nil, result.Error
	}

	return &org, nil
}
