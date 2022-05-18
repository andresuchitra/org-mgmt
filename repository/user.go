package repository

import (
	"context"

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
	GetMembersByOrganizationID(ctx context.Context, organizationID int64) ([]models.UserWithFollower, error)
}

func (m *Repo) GetMembersByOrganizationID(ctx context.Context, organizationID int64) ([]models.UserWithFollower, error) {
	members := make([]models.UserWithFollower, 0)

	result := m.Conn.Raw(`SELECT users.id, users.name, users.email, users.password, users.avatar, users.organization_id,
		following_table.total_following, count(uf.follower_id) AS total_follower
		FROM users LEFT JOIN (
			SELECT u.id, count(*) as total_following FROM users u
			LEFT JOIN user_followers ON u.id = user_followers.follower_id
			WHERE u.organization_id = 1  GROUP BY u.id
		) as following_table ON users.id = following_table.id
		LEFT JOIN user_followers uf ON uf.user_id = users.id
		WHERE organization_id = 1
		GROUP BY users.id, following_table.total_following
		ORDER BY total_follower DESC`).
		Scan(&members)
	if result.Error != nil {
		return nil, result.Error
	}

	return members, nil
}
