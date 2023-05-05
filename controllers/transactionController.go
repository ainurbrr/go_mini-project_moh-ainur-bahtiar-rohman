package controllers

import (
	"net/http"
	"penggalangan-dana/formatter"
	"penggalangan-dana/helpers"
	"penggalangan-dana/lib/database"
	"penggalangan-dana/models"

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
