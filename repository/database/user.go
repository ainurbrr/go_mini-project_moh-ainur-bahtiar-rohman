package database

import (
	"struktur-penggalangan-dana/config"
	"struktur-penggalangan-dana/models"
)

func GetUserById(id int) (*models.User, error) {
	var user models.User

	if err := config.DB.Where("id = ?", user.ID).First(&user).Error; err != nil {
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
