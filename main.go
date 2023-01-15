package main

import (
	"fmt"
	"log"
	"os"
	"testjavan/controllers"
	"testjavan/helpers"
	"testjavan/helpers/constants"
	"testjavan/pkg"
	"testjavan/repositories"
	"testjavan/tcp"
	usecase "testjavan/usecases"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"github.com/subosito/gotenv"
)

func init() {
	if err := gotenv.Load(*helpers.ProjectFolder + ".env"); err == nil {
		fmt.Printf("[%s] Env loaded\n", time.Now().Format(constants.TimeFormat))
	}

	logrus.SetFormatter(&logrus.JSONFormatter{
		DisableTimestamp: true,
	})
}

func main() {
	e := echo.New()

	skipURL := []string{
		"/health-check",
	}

	helpers.LoggerMiddleware(e, skipURL)
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.GET("/health-check", func(c echo.Context) error {
		return c.JSON(200, "OK")
	})

	e.Use(helpers.ValidateContentType())

	db, err := pkg.ConnectDB()
	if err != nil {
		log.Fatal("Failed connecting to database")
	}

	repo := repositories.NewRepository(db)
	usecase := usecase.NewUsecase(repo)
	controllers.NewController(e, usecase)

	go tcp.TCP()

	fmt.Printf("[%s] HTTP Service running on port: %s\n", time.Now().Format(constants.TimeFormat), os.Getenv("APP_PORT"))
	e.HideBanner = true
	e.HidePort = true

	e.Logger.Fatal(e.Start("localhost:" + os.Getenv("APP_PORT")))
}
