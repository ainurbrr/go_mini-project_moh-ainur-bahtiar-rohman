package database

import (
	"struktur-penggalangan-dana/config"
	"struktur-penggalangan-dana/models"
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
