package repository

import (
	"context"
	"log"

	"github.com/andresuchitra/org-mgmt/models"
	"gorm.io/gorm"
)

type followingStruct struct {
	ID             uint32
	TotalFollowing uint32
}

func NewUserRepository(Conn *gorm.DB) UserRepository {
	return &Repo{Conn}
}

type UserRepository interface {
	GetMembersByOrganizationID(ctx context.Context, organizationID int64) ([]models.User, error)
}

func (m *Repo) GetMembersByOrganizationID(ctx context.Context, organizationID int64) ([]models.User, error) {
	members := make([]models.User, 0)
	followings := make([]followingStruct, 0)

	result := m.Conn.Where("organization_id = ?", organizationID).Preload("Followers").Find(&members)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, item := range members {
		log.Println("ID, total followers | ", item.ID, len(item.Followers))
	}

	result = m.Conn.Raw(`SELECT users.id, count(*) as total_following
		FROM "users" LEFT JOIN user_followers ON users.id = user_followers.follower_id
		WHERE organization_id = 1  GROUP BY "id";`).Scan(&followings)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, item := range members {
		log.Println("ID, total Following | ", item.ID, item.TotalFollowing)
	}

	return members, nil
}
