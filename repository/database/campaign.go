package database

import (
	config "github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/config"
	models "github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/models"
)

func FindAllCampaign() ([]models.Campaign, error) {
	var campaigns []models.Campaign

	if err := config.DB.Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error; err != nil {
		return campaigns, err
	}
	return campaigns, nil

}

func FindCampaignByUserId(userId int) ([]models.Campaign, error) {
	var campaigns []models.Campaign

	if err := config.DB.Where("user_id = ?", userId).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error; err != nil {
		return nil, err
	}
	return campaigns, nil
}

func FindCampaignById(id int) (models.Campaign, error) {
	var campaign models.Campaign

	if err := config.DB.Preload("User").Preload("CampaignImages").Where("id = ?", id).Find(&campaign).Error; err != nil {
		return campaign, err
	}
	return campaign, nil
}

func CreateCampaign(campaign models.Campaign) (err error) {
	if err = config.DB.Create(&campaign).Error; err != nil {
		return
	}
	return nil
}

func FindCampaignBySlug(slug string) (models.Campaign, error) {
	var campaign models.Campaign

	if err := config.DB.Preload("User").Preload("CampaignImages").Where("slug = ?", slug).Find(&campaign).Error; err != nil {
		return campaign, err
	}
	return campaign, nil
}

func UpdateCampaign(campaign *models.Campaign) (err error) {
	if err = config.DB.Model(&campaign).Updates(campaign).Error; err != nil {
		return
	}

	return nil
}

func UploadCampaignImage(campaignImage models.Campaign_image) (err error) {
	if err = config.DB.Create(&campaignImage).Error; err != nil {
		return
	}
	return nil
}
