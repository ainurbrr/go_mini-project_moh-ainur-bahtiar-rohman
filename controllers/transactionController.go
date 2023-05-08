package controllers

import (
	"net/http"
	"strconv"

	"github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/formatter"
	"github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/helpers"
	"github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/models/payload"
	"github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/usecase"

	"github.com/labstack/echo/v4"
)

func GetCampaignTransactionsController(c echo.Context) error {
	campaignId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	transactions, err := usecase.GetTransactionsByCampaignId(campaignId, c)
	if err != nil {
		return err
	}

	formatCampaignTransaction := formatter.FormatCampaignTransactions(transactions)
	response := helpers.APIResponse(http.StatusOK, "succes", formatCampaignTransaction, "Successfully get campaign transactions")

	return c.JSON(http.StatusOK, response)

}

func GetUserTransactionsController(c echo.Context) error {
	transactions, err := usecase.GetTransactionByUserId(c)
	if err != nil {
		return err
	}

	formatUserTransaction := formatter.FormatUserTransactions(transactions)
	response := helpers.APIResponse(http.StatusOK, "succes", formatUserTransaction, "Successfully get user transactions")

	return c.JSON(http.StatusOK, response)
}

func CreateTransactionController(c echo.Context) error {

	payloadTransaction := payload.CreateTransactionRequest{}
	c.Bind(&payloadTransaction)

	if err := c.Validate(payloadTransaction); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "error payload create transaction",
			"error":    err.Error(),
		})
	}

	transaction, err := usecase.CreateTransaction(c, &payloadTransaction)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	formatTransaction := formatter.FormatTransaction(transaction)
	response := helpers.APIResponse(http.StatusOK, "succes", formatTransaction, "Successfully created transaction")

	return c.JSON(http.StatusOK, response)
}

// func GetNotificationController(c echo.Context) error {
// 	var input payment.PaymentNotificationInput

// 	err := c.Bind(&input)
// 	if err != nil {
// 		return err
// 	}

// 	err = database.ProcessPayment(c, input)
// 	if err != nil {
// 		return err
// 	}

// 	return c.JSON(http.StatusOK, input)
// }
