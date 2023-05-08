package database

import (

	"github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/config"
	"github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/models"
)

func FindUserById(id int) (*models.User, error) {
	user := models.User{}

	if err := config.DB.Model(&user).Where("id = ?", id).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(user *models.User) error {
	if err := config.DB.Model(&user).Updates(user).Error; err != nil {
		return err
	}
	return nil
}
