package controllers

import (
	"net/http"
	"struktur-penggalangan-dana/formatter"
	"struktur-penggalangan-dana/helpers"
	"struktur-penggalangan-dana/repository/database"
	"struktur-penggalangan-dana/models"
	"struktur-penggalangan-dana/payment"

	"github.com/labstack/echo/v4"
)

func GetCampaignTransactionsController(c echo.Context) error {

	transactions, err := database.GetTransactionsByCampaignId(c)
	if err != nil {
		return err
	}

	campaignTransactionsStruct := transactions.([]models.Transaction)
	formatCampaignTransaction := formatter.FormatCampaignTransactions(campaignTransactionsStruct)
	response := helpers.APIResponse(http.StatusOK, "succes", formatCampaignTransaction, "Successfully get campaign transactions")

	return c.JSON(http.StatusOK, response)

}

func GetUserTransactionsController(c echo.Context) error {
	transactions, err := database.GetTransactionByUserId(c)
	if err != nil {
		return err
	}

	userTransactionsStruct := transactions.([]models.Transaction)
	formatUserTransaction := formatter.FormatUserTransactions(userTransactionsStruct)
	response := helpers.APIResponse(http.StatusOK, "succes", formatUserTransaction, "Successfully get user transactions")

	return c.JSON(http.StatusOK, response)
}

func CreateTransactionController(c echo.Context) error {
	transaction, err := database.CreateTransaction(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	transactionStruct := transaction.(models.Transaction)
	formatTransaction := formatter.FormatTransaction(transactionStruct)
	response := helpers.APIResponse(http.StatusOK, "succes", formatTransaction, "Successfully created transaction")

	return c.JSON(http.StatusOK, response)
}

func GetNotificationController(c echo.Context) error {
	var input payment.PaymentNotificationInput

	err := c.Bind(&input)
	if err != nil {
		return err
	}

	err = database.ProcessPayment(c, input)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, input)
}
