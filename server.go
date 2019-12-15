package main

import (
	"github.com/labstack/echo"
	"github.com/rfdnxbro/youtube-manager-go/routes"
)

func main() {
	e := echo.New()

  // Routes
  routes.Init(e)

  // Start server
	e.Logger.Fatal(e.Start(":8080"))
}
