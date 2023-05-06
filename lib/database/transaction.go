package database

import (
	"errors"
	"penggalangan-dana/config"
	"penggalangan-dana/middlewares"
	"penggalangan-dana/models"
	"penggalangan-dana/payment"
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

func GetByUserId(userId int) (interface{}, error) {
	var transactions []models.Transaction
	if err := config.DB.Preload("User").Preload("Campaign").Preload("Campaign.CampaignImages", "campaign_images.is_primary = 1").Where("user_id = ?", userId).Order("id desc").Find(&transactions).Error; err != nil {
		return transactions, err
	}
	return transactions, nil
}

func GetTransactionByUserId(c echo.Context) (interface{}, error) {
	idFromToken, _ := middlewares.ExtractTokenId(c)
	transactions, err := GetByUserId(idFromToken)
	if err != nil {
		return transactions, err
	}
	return transactions, nil

}

// func SaveTransaction(transaction models.Transaction) (interface{}, error) {
// 	if err := config.DB.Create(&transaction).Error; err != nil {
// 		return transaction, err
// 	}
// 	return transaction, nil
// }

func CreateTransaction(c echo.Context) (interface{}, error) {
	transaction := models.Transaction{}
	c.Bind(&transaction)
	campaign_id, _ := strconv.Atoi(c.FormValue("campaign_id"))
	amount, _ := strconv.Atoi(c.FormValue("amount"))
	transaction.CampaignID = campaign_id //CampaignID
	transaction.Amount = amount
	transaction.Status = "pending"
	transaction.Code = ""

	idFromToken, _ := middlewares.ExtractTokenId(c)
	user, _ := GetById(idFromToken)
	userModel := user.(models.User)
	transaction.User = &userModel

	if err := config.DB.Create(&transaction).Error; err != nil {
		return nil, err
	}

	paymentURL, err := payment.GetPaymentURL(transaction, userModel)
	if err != nil {
		return transaction, err
	}

	transaction.PaymentURL = paymentURL
	transaction, err = UpdateTransaction(transaction)
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func UpdateTransaction(transaction models.Transaction) (models.Transaction, error) {
	if err := config.DB.Save(&transaction).Error; err != nil {
		return transaction, err
	}
	return transaction, nil
}

func GetTransactionById(Id int) (interface{}, error) {
	var transactions models.Transaction
	if err := config.DB.Where("id=?", Id).Error; err != nil {
		return transactions, err
	}
	return transactions, nil

}

func ProcessPayment(c echo.Context, input payment.PaymentNotificationInput) error {
	transactionId, _ := strconv.Atoi(input.OrderID)
	transaction, err := GetTransactionById(transactionId)
	if err != nil {
		return err
	}
	transactionModel := transaction.(models.Transaction)

	if input.PaymentType == "credit_card" && input.TransactionStatus == "camptured" && input.FraudStatus == "accept" {
		transactionModel.Status = "paid"
	} else if input.TransactionStatus == "settlement" {
		transactionModel.Status = "paid"
	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel" {
		transactionModel.Status = "cancelled"
	}

	updatedTransaction, err := UpdateTransaction(transactionModel)
	if err != nil {
		return err
	}

	campaign, err := FindById(updatedTransaction.CampaignID)
	campaignModel := campaign.(models.Campaign)
	if err != nil {
		return err
	}

	if updatedTransaction.Status == "paid" {
		campaignModel.BackerCount = campaignModel.BackerCount + 1
		campaignModel.TotalAmount = campaignModel.TotalAmount + updatedTransaction.Amount

		_, err := UpdateCampaign(c)
		if err != nil {
			return err
		}
	}

	return nil

}
