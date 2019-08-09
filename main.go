package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {

	// load in .env file if it exists
	if fileExists(".env") {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		} else {
			log.Println("Successfully loaded environment variables")
		}
	}

	// initialize the mail client
	initMailClient()

	// initialize cron
	initCron()

	e := echo.New()

	e.POST("/report-uri/:domain", handleReport)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "This is a csp report handler. See https://git.bn4t.me/bn4t/csp-handler for more info.")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
