package usecase

import (
	"errors"
	"penggalangan-dana/config"
	"penggalangan-dana/models"
)

func IsEmailAvailable(email string) bool {
	var count int64
	user := models.User{}
	config.DB.Model(&user).Where("email = ?", email).Count(&count)
	return count == 0
}

func FindById(id int) (interface{}, error) {

	user := models.User{}
	err := config.DB.Model(&user).Where("id = ?", id).Find(&user).Error
	if err != nil {
		return nil, err
	}

	if user.ID == 0 {
		return nil, errors.New("user not found")
	}
	return user, nil
}

