package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"youtube-manager-go/routes"
  "github.com/sirupsen/logrus"
)

func init() {
  logrus.SetLevel(logrus.DebugLevel)
  logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	e := echo.New()

	// Middlewares
	e.Use(middleware.Logger())

	// Routes
	routes.Init(e)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
