package database

import (
	"penggalangan-dana/config"
	"penggalangan-dana/models"
)

func FindAll() (interface{}, error) {
	var campaigns []models.Campaign

	if err := config.DB.Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error; err != nil {
		return campaigns, err
	}
	return campaigns, nil

}

func FindByUserId(user_id int) (interface{}, error) {
	var campaigns []models.Campaign

	if err := config.DB.Where("user_id = ?", user_id).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error; err != nil {
		return campaigns, err
	}
	return campaigns, nil
}
