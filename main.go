package main

import (
	"github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/routes"

	config "github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/config"

	"github.com/labstack/echo/v4"
)

func main() {

	db := config.Init()
	e := echo.New()
	routes.Routes(e, db)

	e.Logger.Fatal(e.Start(":8080"))
}
