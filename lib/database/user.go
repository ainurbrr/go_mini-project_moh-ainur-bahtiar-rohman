package database

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"penggalangan-dana/config"
	"penggalangan-dana/middlewares"
	"penggalangan-dana/models"
	"penggalangan-dana/usecase"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c echo.Context) (interface{}, error) {
	user := models.User{}
	c.Bind(&user)
	user.Role = "user"
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(passwordHash)

	if !usecase.IsEmailAvailable(user.Email) {
		return nil, errors.New("email is already taken")
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func LoginUser(c echo.Context) (interface{}, error) {
	user := models.User{}
	c.Bind(&user)
	password := c.FormValue("password")

	err := config.DB.Where("email = ?", user.Email).First(&user).Error
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	return user, nil
}

func Update(c echo.Context) (interface{}, error) {
	id, err := middlewares.ExtractTokenId(c)
	if err != nil {
		return nil, err
	}

	user, err := usecase.FindById(id)
	if err != nil {
		return nil, err
	}
	userModel := user.(models.User)

	c.Bind(&userModel)
	file, err := c.FormFile("avatar_file_name")
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("images/avatar/%d-%s", userModel.ID, file.Filename)
	userModel.Avatar_File_Name = path

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

	//save to db
	if err := config.DB.Model(&userModel).Updates(userModel).Error; err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return userModel, nil
}
