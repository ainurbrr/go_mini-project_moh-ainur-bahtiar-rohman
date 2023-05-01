package controllers

import (
	"net/http"
	"penggalangan-dana/formatter"
	"penggalangan-dana/helpers"
	"penggalangan-dana/lib/database"
	"penggalangan-dana/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetCampaignController(c echo.Context) error {
	user_id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return err
	}

	campaign, err := database.GetCampaigns(user_id)
	if err != nil {
		return err
	}
	campaignStruct := campaign.([]models.Campaign)
	formatCampaign := formatter.FormatCampaigns(campaignStruct)
	response := helpers.APIResponse(http.StatusOK, "succes", formatCampaign, "Successfully Get Campaigns By User Id")

	return c.JSON(http.StatusOK, response)
}

func GetCampaignsController(c echo.Context) error {
	campaigns, err := database.GetCampaigns(0)
	if err != nil {
		return err
	}
	campaignsStruct := campaigns.([]models.Campaign)
	formatCampaigns := formatter.FormatCampaigns(campaignsStruct)
	response := helpers.APIResponse(http.StatusOK, "succes", formatCampaigns, "Successfully Get All List Campaigns")

	return c.JSON(http.StatusOK, response)
}
