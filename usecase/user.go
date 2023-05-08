package usecase

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/repository/database"

	"github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/middlewares"

	"github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/models"
	"github.com/labstack/echo/v4"
)

func UploadAvatar(c echo.Context) (user *models.User, err error) {
	id, err := middlewares.ExtractTokenId(c)
	if err != nil {
		return
	}

	user, err = database.FindUserById(id)
	if err != nil {
		return nil, err
	}

	file, err := c.FormFile("avatar_file_name")
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("images/avatar/%d-%s", user.ID, file.Filename)
	user.Avatar_File_Name = path

	//upload the avatar
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()
	// Create a new file on disk
	dst, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	defer dst.Close()
	// Copy the uploaded file to the destination file
	if _, err = io.Copy(dst, src); err != nil {
		return nil, err
	}

	if err := database.UpdateUser(user); err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return user, nil
}
