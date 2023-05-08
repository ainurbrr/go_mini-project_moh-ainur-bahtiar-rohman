package controllers

import (
	"net/http"
	"strconv"

	"github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/formatter"
	"github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/models"
	"github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/models/payload"
	"github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/helpers"
	"github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/usecase"

	"github.com/labstack/echo/v4"
)

func GetCampaignsController(c echo.Context) error {

	user_id, _ := strconv.Atoi(c.QueryParam("user_id"))

	campaign, err := usecase.GetCampaigns(user_id)
	if err != nil {
		return err
	}
	campaignModel := campaign.([]models.Campaign)
	formatCampaign := formatter.FormatCampaigns(campaignModel)
	response := helpers.APIResponse(http.StatusOK, "succes", formatCampaign, "Successfully Get Campaigns")

	return c.JSON(http.StatusOK, response)
}

func GetCampaignDetailController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	campaign, err := usecase.GetCampaign(id)
	if err != nil {
		return err
	}
	formatCampaign, err := formatter.FormatCampaignDetail(campaign)
	if err != nil {
		return err
	}
	response := helpers.APIResponse(http.StatusOK, "succes", formatCampaign, "Successfully Get Campaign detail By Id")

	return c.JSON(http.StatusOK, response)
}

func CreateCampaignController(c echo.Context) error {
	payloadCampaign := payload.CreateCampaignRequest{}
	c.Bind(&payloadCampaign)

	if err := c.Validate(payloadCampaign); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "error payload create campaign",
			"error":    err.Error(),
		})
	}
	campaign, err := usecase.CreateCampaign(c, &payloadCampaign)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": "error create campaign",
			"error":    err.Error(),
		})
	}

	formatCampaign := formatter.FormatCampaign(campaign)
	response := helpers.APIResponse(http.StatusOK, "succes", formatCampaign, "Successfully created campaign")

	return c.JSON(http.StatusOK, response)
}

func UpdateCampaignController(c echo.Context) error {
	campaign, err := usecase.UpdateCampaign(c)
	if err != nil {
		return err
	}
	formatCampaignUpdated, _ := formatter.FormatCampaignDetail(campaign)
	response := helpers.APIResponse(http.StatusOK, "succes", formatCampaignUpdated, "Success to Update Campaign")

	return c.JSON(http.StatusOK, response)
}

func UploadCampaignImageController(c echo.Context) error {
	campaignImage, err := usecase.UploadCampaignImage(c)
	if err != nil {
		return err
	}
	response := helpers.APIResponse(http.StatusOK, "succes", campaignImage, "Campaign Image Successfully Uploaded")
	return c.JSON(http.StatusOK, response)
}
