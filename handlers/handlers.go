package handlers

import (
	"github.com/andresuchitra/org-mgmt/repository"
	"github.com/labstack/echo"
)

type Handlers struct {
	CommentRepo      repository.CommentRepository
	UserRepo         repository.UserRepository
	OrganizationRepo repository.OrganizationRepository
}

func NewHandlers(e *echo.Echo, commentRepo repository.CommentRepository,
	userRepo repository.UserRepository, orgRepo repository.OrganizationRepository) {
	handler := &Handlers{
		CommentRepo:      commentRepo,
		UserRepo:         userRepo,
		OrganizationRepo: orgRepo,
	}

	e.GET("/orgs/:id/comments", handler.FetchCommentsByOrganizationID)
	e.GET("/orgs/:id/members", handler.FetchMembersByOrganizationID)

	// e.POST("/orgs/:id/comments", handler.CreateCommentByOrganizationID)
	// e.DELETE("/orgs/:id/comments", handler.DeleteCommentsByOrganizationID)
}
