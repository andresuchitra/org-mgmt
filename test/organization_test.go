package test

import (
	"context"
	"errors"
	"testing"

	"github.com/andresuchitra/org-mgmt/models"
	"github.com/andresuchitra/org-mgmt/repository/mocks"
	"github.com/andresuchitra/org-mgmt/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	organization models.Organization
)

func TestFetchCommentsByOrganizationName(t *testing.T) {
	comments := []models.Comment{
		{
			Text:           "comment 1",
			IsDeleted:      false,
			AuthorID:       1,
			OrganizationID: 1,
		},
		{
			Text:           "comment 2",
			IsDeleted:      false,
			AuthorID:       1,
			OrganizationID: 1,
		},
	}
	organization = models.Organization{
		Name: "xendit",
	}

	t.Run("Succes Return members list", func(t *testing.T) {
		mockOrganizationRepo := mocks.OrganizationRepository{}
		mockCommentRepo := mocks.CommentRepository{}
		mockUserRepo := mocks.UserRepository{}
		ctx := context.Background()

		mockOrganizationRepo.On("GetOrganizationByName", mock.Anything, "xendit").Return(&organization, nil)
		mockCommentRepo.On("GetCommentsByOrgID", mock.Anything, organization.ID).Return(comments, nil)
		orgService := service.NewOrganizationService(&mockOrganizationRepo, &mockUserRepo, &mockCommentRepo)

		result, err := orgService.FetchCommentsByOrganizationName(ctx, "xendit")

		assert.Nil(t, err)
		assert.Equal(t, len(comments), len(result))
		for i, item := range result {
			assert.IsType(t, models.CommentResponse{}, item)
			assert.Equal(t, comments[i].Text, item.Text)
		}
	})

	t.Run("Fail Error fetching organization", func(t *testing.T) {
		mockOrganizationRepo := mocks.OrganizationRepository{}
		mockCommentRepo := mocks.CommentRepository{}
		mockUserRepo := mocks.UserRepository{}
		ctx := context.Background()
		mockOrganizationRepo.On("GetOrganizationByName", mock.Anything, "xendit").Return(nil, errors.New("organization not found"))
		mockCommentRepo.On("GetCommentsByOrgID", mock.Anything, organization.ID).Return(nil, nil)

		orgService := service.NewOrganizationService(&mockOrganizationRepo, &mockUserRepo, &mockCommentRepo)

		result, err := orgService.FetchCommentsByOrganizationName(ctx, "xendit")

		assert.Nil(t, result)
		assert.Equal(t, "organization not found", err.Error())
	})
}

func TestFetchMembersByOrganizationName(t *testing.T) {
	users := []models.User{
		{
			Name:           "first",
			OrganizationID: 1,
		},
		{
			Name:           "second",
			OrganizationID: 1,
		},
	}

	members := []models.UserWithFollower{
		{
			User:           users[0],
			TotalFollowing: 1,
			TotalFollower:  1,
		},
		{
			User:           users[1],
			TotalFollowing: 1,
			TotalFollower:  1,
		},
	}
	organization = models.Organization{
		Name: "xendit",
	}

	t.Run("Succes Return comments", func(t *testing.T) {
		mockOrganizationRepo := mocks.OrganizationRepository{}
		mockCommentRepo := mocks.CommentRepository{}
		mockUserRepo := mocks.UserRepository{}
		ctx := context.Background()

		mockOrganizationRepo.On("GetOrganizationByName", mock.Anything, "xendit").Return(&organization, nil)
		mockUserRepo.On("GetMembersByOrganizationID", mock.Anything, int64(organization.ID)).Return(members, nil)
		orgService := service.NewOrganizationService(&mockOrganizationRepo, &mockUserRepo, &mockCommentRepo)

		result, err := orgService.FetchMembersByOrganizationName(ctx, "xendit")

		assert.Nil(t, err)
		assert.Equal(t, len(members), len(result))
		for i, item := range result {
			assert.IsType(t, models.UserResponse{}, item)
			assert.Equal(t, members[i].Name, item.Name)
		}
	})

	t.Run("Fail Error fetching organization", func(t *testing.T) {
		mockOrganizationRepo := mocks.OrganizationRepository{}
		mockCommentRepo := mocks.CommentRepository{}
		mockUserRepo := mocks.UserRepository{}
		ctx := context.Background()
		mockOrganizationRepo.On("GetOrganizationByName", mock.Anything, "xendit").Return(nil, errors.New("organization not found"))
		mockUserRepo.On("GetMembersByOrganizationID", mock.Anything, int64(organization.ID)).Return(members, nil)

		orgService := service.NewOrganizationService(&mockOrganizationRepo, &mockUserRepo, &mockCommentRepo)

		result, err := orgService.FetchMembersByOrganizationName(ctx, "xendit")

		assert.Nil(t, result)
		assert.Equal(t, "organization not found", err.Error())
	})
}

func TestCreateCommentByOrganizationName(t *testing.T) {
	organization = models.Organization{
		Name: "xendit",
	}
	param := models.CreateCommentParam{
		OrganizationName: organization.Name,
		Comment:          "new comment",
		AuthorID:         1,
	}

	t.Run("Succes creating comment", func(t *testing.T) {
		mockOrganizationRepo := mocks.OrganizationRepository{}
		mockCommentRepo := mocks.CommentRepository{}
		mockUserRepo := mocks.UserRepository{}
		ctx := context.Background()

		mockOrganizationRepo.On("GetOrganizationByName", mock.Anything, "xendit").Return(&organization, nil)
		mockCommentRepo.On("CreateCommentByOrganizationID", mock.Anything, uint(organization.ID), param.AuthorID, param.Comment).Return(nil)
		orgService := service.NewOrganizationService(&mockOrganizationRepo, &mockUserRepo, &mockCommentRepo)

		err := orgService.CreateCommentByOrganizationName(ctx, param)

		assert.Nil(t, err)
	})

	t.Run("Fail creating comment", func(t *testing.T) {
		mockOrganizationRepo := mocks.OrganizationRepository{}
		mockCommentRepo := mocks.CommentRepository{}
		mockUserRepo := mocks.UserRepository{}
		ctx := context.Background()
		mockOrganizationRepo.On("GetOrganizationByName", mock.Anything, "xendit").Return(nil, errors.New("organization not found"))
		mockCommentRepo.On("CreateCommentByOrganizationID", mock.Anything, uint(organization.ID), param.AuthorID, param.Comment).Return(nil)

		orgService := service.NewOrganizationService(&mockOrganizationRepo, &mockUserRepo, &mockCommentRepo)

		err := orgService.CreateCommentByOrganizationName(ctx, param)

		assert.NotNil(t, err)
		assert.Equal(t, "organization not found", err.Error())
	})

	t.Run("Fail - invalid author", func(t *testing.T) {
		mockOrganizationRepo := mocks.OrganizationRepository{}
		mockCommentRepo := mocks.CommentRepository{}
		mockUserRepo := mocks.UserRepository{}
		ctx := context.Background()

		mockOrganizationRepo.On("GetOrganizationByName", mock.Anything, "xendit").Return(&organization, nil)
		mockCommentRepo.On("CreateCommentByOrganizationID", mock.Anything, uint(organization.ID), param.AuthorID, param.Comment).Return(errors.New("Invalid Author"))

		orgService := service.NewOrganizationService(&mockOrganizationRepo, &mockUserRepo, &mockCommentRepo)

		err := orgService.CreateCommentByOrganizationName(ctx, param)

		assert.NotNil(t, err)
		assert.Equal(t, "Invalid Author", err.Error())
	})
}

func TestSoftDeleteCommentsByOrganizationName(t *testing.T) {
	organization = models.Organization{
		Name: "xendit",
	}

	t.Run("Succes soft delete comments", func(t *testing.T) {
		mockOrganizationRepo := mocks.OrganizationRepository{}
		mockCommentRepo := mocks.CommentRepository{}
		mockUserRepo := mocks.UserRepository{}
		ctx := context.Background()

		mockOrganizationRepo.On("GetOrganizationByName", mock.Anything, "xendit").Return(&organization, nil)
		mockCommentRepo.On("SoftDeleteCommentsByOrganizationID", mock.Anything, uint(organization.ID)).Return(nil)
		orgService := service.NewOrganizationService(&mockOrganizationRepo, &mockUserRepo, &mockCommentRepo)

		err := orgService.SoftDeleteCommentsByOrganizationName(ctx, organization.Name)

		assert.Nil(t, err)
	})

	t.Run("Fail creating comment", func(t *testing.T) {
		mockOrganizationRepo := mocks.OrganizationRepository{}
		mockCommentRepo := mocks.CommentRepository{}
		mockUserRepo := mocks.UserRepository{}
		ctx := context.Background()

		mockOrganizationRepo.On("GetOrganizationByName", mock.Anything, "xendit").Return(&organization, nil)
		mockCommentRepo.On("SoftDeleteCommentsByOrganizationID", mock.Anything, uint(organization.ID)).Return(errors.New("organization not found"))
		orgService := service.NewOrganizationService(&mockOrganizationRepo, &mockUserRepo, &mockCommentRepo)

		err := orgService.SoftDeleteCommentsByOrganizationName(ctx, organization.Name)

		assert.NotNil(t, err)
		assert.Equal(t, "organization not found", err.Error())
	})
}
