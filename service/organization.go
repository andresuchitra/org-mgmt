package service

import (
	"context"

	"github.com/andresuchitra/org-mgmt/models"
	"github.com/andresuchitra/org-mgmt/repository"
)

type OrganizationService interface {
	FetchCommentsByOrganizationName(ctx context.Context, name string) ([]models.CommentResponse, error)
	FetchMembersByOrganizationName(ctx context.Context, name string) ([]models.UserResponse, error)
	CreateCommentByOrganizationName(context.Context, models.CreateCommentParam) error
	SoftDeleteCommentsByOrganizationName(ctx context.Context, name string) error
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

func (s *organizationService) FetchCommentsByOrganizationName(ctx context.Context, name string) ([]models.CommentResponse, error) {
	// 1. get organization ID by name
	org, err := s.orgRepo.GetOrganizationByName(ctx, name)
	if err != nil {
		return nil, err
	}

	// 2. call comment repo to get comments by org ID
	comments, err := s.commentRepo.GetCommentsByOrgID(ctx, org.ID)
	if err != nil {
		return nil, err
	}

	commentResponses := make([]models.CommentResponse, 0)
	for _, item := range comments {
		commentResponses = append(commentResponses, item.ConvertToCommentResponse())
	}

	return commentResponses, nil
}

func (s *organizationService) FetchMembersByOrganizationName(ctx context.Context, name string) ([]models.UserResponse, error) {
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

	membersResponse := make([]models.UserResponse, 0)
	for _, item := range members {
		newUser := item.ConvertToUserResponse()
		membersResponse = append(membersResponse, newUser)
	}

	return membersResponse, nil
}

func (s *organizationService) CreateCommentByOrganizationName(ctx context.Context, param models.CreateCommentParam) error {
	org, err := s.orgRepo.GetOrganizationByName(ctx, param.OrganizationName)
	if err != nil {
		return err
	}

	err = s.commentRepo.CreateCommentByOrganizationID(ctx, org.ID, param.AuthorID, param.Comment)
	if err != nil {
		return err
	}

	return nil
}

func (s *organizationService) SoftDeleteCommentsByOrganizationName(ctx context.Context, name string) error {
	org, err := s.orgRepo.GetOrganizationByName(ctx, name)
	if err != nil {
		return err
	}

	err = s.commentRepo.SoftDeleteCommentsByOrganizationID(ctx, org.ID)
	if err != nil {
		return err
	}

	return nil
}
