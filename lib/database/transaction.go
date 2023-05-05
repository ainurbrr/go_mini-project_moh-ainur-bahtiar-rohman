package database

import (
	"errors"
	"penggalangan-dana/config"
	"penggalangan-dana/middlewares"
	"penggalangan-dana/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func FindByCampaignId(campaignId int) (interface{}, error) {
	var transactions []models.Transaction

	if err := config.DB.Preload("User").Preload("Campaign").Where("campaign_id = ?", campaignId).Order("id desc").Find(&transactions).Error; err != nil {
		return transactions, err
	}
	return transactions, nil
}

func GetTransactionsByCampaignId(c echo.Context) (interface{}, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return nil, err
	}
	campaign, _ := FindById(id)
	idFromToken, err := middlewares.ExtractTokenId(c)
	if err != nil {
		return nil, err
	}
	transactions, err := FindByCampaignId(id)
	if err != nil {
		return nil, err
	}

	campaignModel := campaign.(models.Campaign)

	if campaignModel.UserID != idFromToken {
		return nil, errors.New("Unauthorized")
	}

	return transactions, nil
}

func GetByUserId(userId int)(interface{}, error){
	var transactions []models.Transaction
	if err := config.DB.Preload("User").Preload("Campaign").Preload("Campaign.CampaignImages", "campaign_images.is_primary = 1").Where("user_id = ?", userId).Order("id desc").Find(&transactions).Error; err != nil{
		return transactions, err
	}
	return transactions, nil
}


func GetTransactionByUserId(c echo.Context)(interface{}, error) {
	idFromToken, _ := middlewares.ExtractTokenId(c)
	transactions, err := GetByUserId(idFromToken)
	if err != nil {
		return transactions, err
	}
	return transactions, nil

}