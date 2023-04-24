package controllers

import (
	"net/http"
	"penggalangan-dana/formatter"
	"penggalangan-dana/lib/database"
	"penggalangan-dana/models"

	"github.com/labstack/echo/v4"
)

func RegisterUserController(c echo.Context) error {
	user, err := database.RegisterUser(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	userStruct := user.(models.User)

	formatUser := formatter.FormatUser(userStruct, "tokenjwt")

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    formatUser,
	})
}
