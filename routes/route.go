package routes

import (
	"penggalangan-dana/middlewares"
	"penggalangan-dana/controllers"

	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func Routes() *echo.Echo {
	e := echo.New()
	middlewares.LogMiddleware(e)
	e.Pre(mid.RemoveTrailingSlash())

	e.POST("/users", controllers.RegisterUserController)

	return e
}
