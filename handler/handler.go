package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/andresuchitra/org-mgmt/models"
	"github.com/andresuchitra/org-mgmt/service"
	echo "github.com/labstack/echo/v4"
)

type Handler struct {
	OrganizationService service.OrganizationService
}

type ResponseError struct {
	Message string `json:"message"`
}

func NewHandler(e *echo.Echo, service service.OrganizationService) {
	handler := &Handler{
		OrganizationService: service,
	}

	e.GET("/orgs/:name/comments", handler.FetchCommentsByOrganizationID)
	e.GET("/orgs/:name/members", handler.FetchMembersByOrganizationID)
	e.POST("/orgs/:name/comments", handler.CreateCommentByOrganizationName)
	e.DELETE("/orgs/:name/comments", handler.SoftDeleteCommentsByOrganizationName)

	// e.POST("/orgs/:id/comments", handler.CreateCommentByOrganizationID)
	// e.DELETE("/orgs/:id/comments", handler.DeleteCommentsByOrganizationID)
}

func (h *Handler) FetchCommentsByOrganizationID(c echo.Context) error {
	orgName := c.Param("name")
	if orgName == "" {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: "Invalid org name"})
	}

	c.Logger().Debug("org name: ", orgName)
	comments, err := h.OrganizationService.FetchCommentsByOrganizationName(c.Request().Context(), strings.ToLower(orgName))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, comments)
}

func (h *Handler) FetchMembersByOrganizationID(c echo.Context) error {
	orgName := c.Param("name")
	if orgName == "" {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: "Invalid org name"})
	}

	members, err := h.OrganizationService.FetchMembersByOrganizationName(c.Request().Context(), strings.ToLower(orgName))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, members)
}

func (h *Handler) CreateCommentByOrganizationName(c echo.Context) error {
	orgName := c.Param("name")
	if orgName == "" {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: "Invalid org name"})
	}

	params := models.CreateCommentParam{}
	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: "Invalid payload format"})
	}

	params.OrganizationName = orgName

	c.Logger().Debug("params: ", params)
	err := h.OrganizationService.CreateCommentByOrganizationName(c.Request().Context(), params)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, fmt.Sprintf("New comment saved to org (%s)!", orgName))
}

func (h *Handler) SoftDeleteCommentsByOrganizationName(c echo.Context) error {
	orgName := c.Param("name")
	if orgName == "" {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: "Invalid org name"})
	}

	err := h.OrganizationService.SoftDeleteCommentsByOrganizationName(c.Request().Context(), strings.ToLower(orgName))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, fmt.Sprintf("All comments for organization (%s) are soft-deleted", orgName))
}
