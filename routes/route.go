package routes

import (
	"penggalangan-dana/constants"
	"penggalangan-dana/controllers"
	"penggalangan-dana/middlewares"

	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func Routes() *echo.Echo {
	e := echo.New()
	middlewares.LogMiddleware(e)
	e.Pre(mid.RemoveTrailingSlash())

	e.POST("/users", controllers.RegisterUserController)
	e.POST("/login", controllers.LoginUserController)
	e.PUT("/avatar", controllers.UploadAvatarController, mid.JWT([]byte(constants.SECRET_JWT)))

	return e
}
