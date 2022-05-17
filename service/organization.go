package service

import (
	"context"
	"log"

	"github.com/andresuchitra/org-mgmt/models"
	"github.com/andresuchitra/org-mgmt/repository"
)

type OrganizationService interface {
	FetchCommentsByOrganizationName(ctx context.Context, name string) ([]models.Comment, error)
	FetchMembersByOrganizationName(ctx context.Context, name string) ([]models.User, error)
}

type organizationService struct {
	orgRepo     repository.OrganizationRepository
	userRepo    repository.UserRepository
	commentRepo repository.CommentRepository
}

func NewOrganizationService(orgRepo repository.OrganizationRepository, userRepo repository.UserRepository, commentRepo repository.CommentRepository) OrganizationService {
	return &organizationService{
		orgRepo:     orgRepo,
		userRepo:    userRepo,
		commentRepo: commentRepo,
	}
}

func (s *organizationService) FetchCommentsByOrganizationName(ctx context.Context, name string) ([]models.Comment, error) {
	// 1. get organization ID by name
	org, err := s.orgRepo.GetOrganizationByName(ctx, name)
	if err != nil {
		return nil, err
	}

	log.Println("org data: ", org)

	// 2. call comment repo to get comments by org ID
	comments, err := s.commentRepo.GetCommentsByOrgID(ctx, int64(org.ID))
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (s *organizationService) FetchMembersByOrganizationName(ctx context.Context, name string) ([]models.User, error) {
	// 1. get organization ID by name
	org, err := s.orgRepo.GetOrganizationByName(ctx, name)
	if err != nil {
		return nil, err
	}

	// 2. call user repo to get users by org ID
	members, err := s.userRepo.GetMembersByOrganizationID(ctx, int64(org.ID))
	if err != nil {
		return nil, err
	}

	return members, nil
}
