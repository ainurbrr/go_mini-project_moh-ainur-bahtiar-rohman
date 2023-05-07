package controllers

import (
	"net/http"
	"struktur-penggalangan-dana/formatter"
	"struktur-penggalangan-dana/helpers"
	"struktur-penggalangan-dana/middlewares"
	"struktur-penggalangan-dana/models/payload"
	"struktur-penggalangan-dana/usecase"

	"github.com/labstack/echo/v4"
)

func LoginUserController(c echo.Context) error {
	var payloadLogin payload.LoginRequest
	c.Bind(&payloadLogin)

	if err := c.Validate(payloadLogin); err != nil {
		return err
	}

	user, err := usecase.LoginUser(c, &payloadLogin)
	c.Bind(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	token, err := middlewares.GenerateToken(user.ID)
	if err != nil {
		return err
	}

	formatUser := formatter.FormatUser(user, token)
	response := helpers.APIResponse(http.StatusOK, "succes", formatUser, "Login Successfully")

	return c.JSON(http.StatusOK, response)
}
