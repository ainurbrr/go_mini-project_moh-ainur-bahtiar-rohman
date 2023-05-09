package database

import (
	"github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/config"
	"github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/models"
)

func FindTransactionByCampaignId(campaignId int) ([]models.Transaction, error) {
	var transactions []models.Transaction

	if err := config.DB.Preload("User").Preload("Campaign").Where("campaign_id = ?", campaignId).Order("id desc").Find(&transactions).Error; err != nil {
		return transactions, err
	}
	return transactions, nil
}

func FindTransactionByUserId(userId int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	if err := config.DB.Preload("User").Preload("Campaign").Preload("Campaign.CampaignImages", "campaign_images.is_primary = 1").Where("user_id = ?", userId).Order("id desc").Find(&transactions).Error; err != nil {
		return transactions, err
	}
	return transactions, nil
}

func CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	if err := config.DB.Create(&transaction).Error; err != nil {
		return transaction, err
	}
	return transaction, nil
}

func UpdateTransaction(transaction models.Transaction) (models.Transaction, error) {
	if err := config.DB.Model(&transaction).Where("id = ?", transaction.ID).Updates(transaction).Error; err != nil {
		return transaction, err
	}
	return transaction, nil
}

func GetTransactionById(Id int) (models.Transaction, error) {
	var transactions models.Transaction
	if err := config.DB.Where("id = ?", Id).Find(&transactions).Error; err != nil {
		return transactions, err
	}
	return transactions, nil

}
