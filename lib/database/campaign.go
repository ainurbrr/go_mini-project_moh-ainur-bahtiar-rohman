package database

import (
	"fmt"
	"penggalangan-dana/config"
	"penggalangan-dana/middlewares"
	"penggalangan-dana/models"
	"time"

	"github.com/gosimple/slug"
	"github.com/labstack/echo/v4"
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

func GetCampaigns(user_id int) (interface{}, error) {
	if user_id != 0 {
		campaigns, err := FindByUserId(user_id)
		if err != nil {
			return campaigns, err
		}
		return campaigns, nil
	}

	campaigns, err := FindAll()
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func FindById(id int) (interface{}, error) {
	var campaign models.Campaign

	if err := config.DB.Preload("User").Preload("CampaignImages").Where("id = ?", id).Find(&campaign).Error; err != nil {
		return nil, err
	}
	return campaign, nil
}

func CreateCampaign(c echo.Context) (interface{}, error) {
	campaign := models.Campaign{}
	c.Bind(&campaign)

	id, err := middlewares.ExtractTokenId(c)
	if err != nil {
		return nil, err
	}
	campaign.UserID = id
	strSlug := fmt.Sprintf("%s %d", campaign.Name, campaign.UserID)
	campaign.Slug = slug.Make(strSlug)
	time, _ := time.Parse("02/01/2006", c.FormValue("end_date"))
	campaign.EndDate = time

	if err := config.DB.Create(&campaign).Error; err != nil {
		return nil, err
	}
	return campaign, nil

}
