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

func NewHandlers(e *echo.Echo, commentRepo repository.CommentRepository) {
	handler := &Handlers{
		Service: us,
	}
	e.GET("/articles", handler.FetchArticle)
	e.POST("/articles", handler.Store)
	e.GET("/articles/:id", handler.GetByID)
	e.DELETE("/articles/:id", handler.Delete)
}
