package main

import (
	"penggalangan-dana/config"
	"penggalangan-dana/routes"
)

func main() {

	config.Init()
	e := routes.Routes()

	e.Logger.Fatal(e.Start(":8080"))
}
