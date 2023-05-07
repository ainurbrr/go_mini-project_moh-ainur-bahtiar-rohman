package main

import (
	"struktur-penggalangan-dana/config"
	"struktur-penggalangan-dana/routes"

	"github.com/labstack/echo/v4"
)

func main() {

	db := config.Init()
	e := echo.New()
	routes.Routes(e, db)

	e.Logger.Fatal(e.Start(":8080"))
}
