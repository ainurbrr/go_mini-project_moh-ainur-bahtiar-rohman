package controllers

import (
	"net/http"
	"penggalangan-dana/formatter"
	"penggalangan-dana/helpers"
	"penggalangan-dana/lib/database"
	"penggalangan-dana/middlewares"
	"penggalangan-dana/models"

	"github.com/labstack/echo/v4"
)

func RegisterUserController(c echo.Context) error {
	user, err := database.RegisterUser(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	userStruct := user.(models.User)

	token, err := middlewares.GenerateToken(userStruct.ID)
	if err != nil {
		return err
	}

	formatUser := formatter.FormatUser(userStruct, token)
	response := helpers.APIResponse(http.StatusOK, "succes", formatUser, "User Registered Successfully")

	return c.JSON(http.StatusOK, response)
}

func LoginUserController(c echo.Context) error {
	userLogin, err := database.LoginUser(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userStruct := userLogin.(models.User)
	token, err := middlewares.GenerateToken(userStruct.ID)
	if err != nil {
		return err
	}

	formatUser := formatter.FormatUser(userStruct, token)
	response := helpers.APIResponse(http.StatusOK, "succes", formatUser, "Login Successfully")

	return c.JSON(http.StatusOK, response)
}

func UploadAvatarController(c echo.Context) error {
	user, err := database.Update(c)
	if err != nil {
		return err
	}
	response := helpers.APIResponse(http.StatusOK, "succes", user, "Avatar Successfully Uploaded")

	return c.JSON(http.StatusOK, response)
}
