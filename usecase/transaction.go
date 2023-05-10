package usecase

import (
	"errors"
	"strconv"

	"github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/models"
	"github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/payment"
	"github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/repository/database"

	"github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/middlewares"
	"github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/models/payload"

	"github.com/labstack/echo/v4"
)

func GetTransactionsByCampaignId(campaign_id int, c echo.Context) (transaction []models.Transaction, err error) {

	campaign, _ := database.FindCampaignById(campaign_id)
	idFromToken, err := middlewares.ExtractTokenId(c)
	if err != nil {
		return
	}
	transaction, err = database.FindTransactionByCampaignId(campaign_id)
	if err != nil {
		return
	}

	if campaign.UserID != idFromToken {
		return transaction, errors.New("Unauthorized")
	}

	return transaction, nil
}

func GetTransactionByUserId(c echo.Context) (transaction []models.Transaction, err error) {
	idFromToken, _ := middlewares.ExtractTokenId(c)
	transactions, err := database.FindTransactionByUserId(idFromToken)
	if err != nil {
		return transactions, err
	}
	return transactions, nil

}

func CreateTransaction(c echo.Context, req *payload.CreateTransactionRequest) (transaction models.Transaction, err error) {

	idFromToken, _ := middlewares.ExtractTokenId(c)

	user, _ := database.FindUserById(idFromToken)

	transaction = models.Transaction{
		CampaignID: req.CampaignID, //CampaignID
		Amount:     req.Amount,
		Status:     "pending",
		Code:       "",
		User:       user,
	}

	transaction, err = database.CreateTransaction(transaction)
	if err != nil {
		return transaction, err
	}

	paymentURL, err := payment.GetPaymentURL(transaction, *user)
	if err != nil {
		return transaction, err
	}

	transaction.PaymentURL = paymentURL
	transactionResult, err := database.UpdateTransaction(transaction)
	if err != nil {
		return transactionResult, err
	}

	return transactionResult, nil
}

func ProcessPayment(c echo.Context, input *payment.PaymentNotificationInput) (*models.Campaign, error) {
	transactionId, _ := strconv.Atoi(input.OrderID)
	transaction, err := database.GetTransactionById(transactionId)
	if err != nil {
		return nil, err
	}
	

	if input.PaymentType == "credit_card" && input.TransactionStatus == "captured" && input.FraudStatus == "accept" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "settlement" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel" {
		transaction.Status = "cancelled"
	}

	updatedTransaction, err := database.UpdateTransaction(transaction)
	if err != nil {
		return nil, err
	}

	campaign, err := database.FindCampaignById(updatedTransaction.CampaignID)

	if err != nil {
		return nil, err
	}

	if updatedTransaction.Status == "paid" {
		campaign.BackerCount = campaign.BackerCount + 1
		campaign.TotalAmount = campaign.TotalAmount + updatedTransaction.Amount

		err := database.UpdateCampaign(&campaign)
		if err != nil {
			return nil, err
		}
	}

	return &campaign, nil
}

// func ProcessPayment(c echo.Context) (*models.Campaign, error) {
// 	transactionId, _ := strconv.Atoi(c.Param("id"))
// 	transaction, err := database.GetTransactionById(transactionId)
// 	if err != nil {
// 		return nil, err
// 	}
// 	c.Bind(&transaction)

// 	if transaction.Status == "paid" {
// 		return nil, errors.New("transaksi sudah dibayar")
// 	}

// 	transaction.Status = "paid"

// 	updatedTransaction, err := database.UpdateTransaction(transaction)
// 	if err != nil {
// 		return nil, err
// 	}

// 	campaign, err := database.FindCampaignById(updatedTransaction.CampaignID)

// 	if err != nil {
// 		return nil, err
// 	}

// 	if updatedTransaction.Status == "paid" {
// 		campaign.BackerCount = campaign.BackerCount + 1
// 		campaign.TotalAmount = campaign.TotalAmount + updatedTransaction.Amount

// 		err := database.UpdateCampaign(&campaign)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}
// 	return &campaign, nil
// }
