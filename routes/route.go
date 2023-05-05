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

	e.Static("/images/avatar", "./images/avatar")
	e.Static("/images/campaignImages", "./images/campaignImages")

	e.POST("/users", controllers.RegisterUserController)
	e.POST("/login", controllers.LoginUserController)
	e.PUT("/avatar", controllers.UploadAvatarController, mid.JWT([]byte(constants.SECRET_JWT)))

	e.GET("/campaigns", controllers.GetCampaignsController)
	e.GET("/campaigns/:id", controllers.GetCampaignController)
	e.POST("/campaign", controllers.CreateCampaignController, mid.JWT([]byte(constants.SECRET_JWT)))
	e.PUT("/campaigns/:id", controllers.UpdateCampaignController, mid.JWT([]byte(constants.SECRET_JWT)))
	e.POST("/campaign-images", controllers.UploadCampaignImageController, mid.JWT([]byte(constants.SECRET_JWT)))

	e.GET("/campaigns/:id/transactions", controllers.GetCampaignTransactionsController, mid.JWT([]byte(constants.SECRET_JWT)))
	e.GET("/transactions", controllers.GetUserTransactionsController, mid.JWT([]byte(constants.SECRET_JWT)))

	return e
}
