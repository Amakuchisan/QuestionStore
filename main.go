package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/Amakuchisan/QuestionStore/database"
	"github.com/Amakuchisan/QuestionStore/route"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/middleware"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
)

func main() {

	if os.Getenv("QS_ENV") == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Could not load .env file")
		}
	}

	e := route.Init()
	e.Use(middleware.Logger())
	e.Static("/static", "static")

	host := os.Getenv("QS_HOST")
	googleCallbackURL := fmt.Sprintf("http://%s/auth/callback/google", host)
	// setup gomniauth
	gomniauth.SetSecurityKey(os.Getenv("SECURITY_KEY"))
	gomniauth.WithProviders(
		google.New(
			os.Getenv("GOOGLE_CLIENT_ID"),
			os.Getenv("GOOGLE_CLIENT_SECRET"),
			googleCallbackURL,
		),
	)

	e.Logger.Fatal(e.Start(":1323"))
}
