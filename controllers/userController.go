package controllers

import (
	"net/http"

	"github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/models/payload"
	"github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/helpers"
	"github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/usecase"

	"github.com/labstack/echo/v4"
)

func RegisterUserController(c echo.Context) error {
	payloadUser := payload.CreateUserRequest{}
	c.Bind(&payloadUser)

	if err := c.Validate(payloadUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "error payload create user",
			"error":    err.Error(),
		})
	}

	resp, err := usecase.CreateUser(&payloadUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": "error create user",
			"error":    err.Error(),
		})
	}

	response := helpers.APIResponse(http.StatusOK, "success", resp, "Succes! Account has been registered")

	return c.JSON(http.StatusOK, response)
}

func UploadAvatarController(c echo.Context) error {
	user, err := usecase.UploadAvatar(c)
	if err != nil {
		return err
	}
	response := helpers.APIResponse(http.StatusOK, "succes", user, "Avatar Successfully Uploaded")

	return c.JSON(http.StatusOK, response)
}
