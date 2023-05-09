package usecase

import (
	"errors"
	"fmt"

	"github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/repository/database"

	"github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/middlewares"
	"github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/models"
	"github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/models/payload"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(req *payload.CreateUserRequest) (resp payload.CreateUserResponse, err error) {

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	if !database.IsEmailAvailable(req.Email) {

		return resp, errors.New("email is already taken")
	}

	newUser := &models.User{
		Name:       req.Name,
		Email:      req.Email,
		Occupation: req.Occupation,
		Password:   string(passwordHash),
		Role:       "user",
	}

	err = database.CreateUser(newUser)
	if err != nil {
		return
	}

	token, err := middlewares.GenerateToken(newUser.ID)
	if err != nil {
		fmt.Println("GetUser: Error generating token")
	}

	resp = payload.CreateUserResponse{
		UserID: newUser.ID,
		Name:   newUser.Name,
		Token:  token,
	}
	return
}

func LoginUser(c echo.Context, req *payload.LoginRequest) (user models.User, err error) {
	user, err = database.LoginUser(req.Email)
	if err != nil {
		fmt.Println("GetUser: Error getting user from database")
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return
	}
	return user, nil

}
