package main

import (
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"github.com/andresuchitra/org-mgmt/handler"
	"github.com/andresuchitra/org-mgmt/pkg/db"
	"github.com/andresuchitra/org-mgmt/repository"
	"github.com/andresuchitra/org-mgmt/service"
)

func main() {
	// init .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:  %s", err.Error())
	}

	DB := db.Init()
	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	e.Logger.Debug(DB)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "welcome")
	})

	orgRepo := repository.NewOrganizationRepository(DB)
	commentRepo := repository.NewCommentRepository(DB)
	userRepo := repository.NewUserRepository(DB)

	orgService := service.NewOrganizationService(orgRepo, userRepo, commentRepo)
	handler.NewHandler(e, orgService)

	// Start server
	log.Fatal(e.Start(":9090"))
}
