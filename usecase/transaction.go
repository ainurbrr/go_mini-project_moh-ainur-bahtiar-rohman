package usecase

import (
	"errors"
	"github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/payment"
	"github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/repository/database"

	middlewares "github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/middlewares"
	models "github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/models"
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
