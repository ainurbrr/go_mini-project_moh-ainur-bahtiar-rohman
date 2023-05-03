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

func GetCampaignsController(c echo.Context) error {
	
	user_id, _ := strconv.Atoi(c.QueryParam("user_id"))

	campaign, err := database.GetCampaigns(user_id)
	if err != nil {
		return err
	}
	campaignStruct := campaign.([]models.Campaign)
	formatCampaign := formatter.FormatCampaigns(campaignStruct)
	response := helpers.APIResponse(http.StatusOK, "succes", formatCampaign, "Successfully Get Campaigns")

	return c.JSON(http.StatusOK, response)
}

func GetCampaignController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	campaign, err := database.FindById(id)
	if err != nil {
		return err
	}
	campaignsStruct := campaign.(models.Campaign)
	formatCampaign, err := formatter.FormatCampaignDetail(campaignsStruct)
	if err != nil {
		return err
	}
	response := helpers.APIResponse(http.StatusOK, "succes", formatCampaign, "Successfully Get Campaign detail By Id")

	return c.JSON(http.StatusOK, response)
}

func CreateCampaignController(c echo.Context) error {
	campaign, err := database.CreateCampaign(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	campaignStruct := campaign.(models.Campaign)
	formatCampaign := formatter.FormatCampaign(campaignStruct)
	response := helpers.APIResponse(http.StatusOK, "succes", formatCampaign, "Successfully created campaign")

	return c.JSON(http.StatusOK, response)
}

func UpdateCampaignController(c echo.Context) error {
	campaign, err := database.UpdateCampaign(c)
	if err != nil {
		return err
	}
	response := helpers.APIResponse(http.StatusOK, "succes", campaign, "Success to Update Campaign")

	return c.JSON(http.StatusOK, response)
}

