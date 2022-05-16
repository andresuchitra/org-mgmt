package main

import (
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"github.com/andresuchitra/org-mgmt/pkg/db"
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

	// Start server
	go func() {
		if err := e.Start(":9090"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()
}
