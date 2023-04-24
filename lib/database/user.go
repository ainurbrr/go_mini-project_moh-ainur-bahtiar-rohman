package database

import (
	"penggalangan-dana/config"
	"penggalangan-dana/models"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c echo.Context) (interface{}, error) {
	user := models.User{}
	c.Bind(&user)

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }
	user.Password = string(passwordHash)

	if err := config.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
