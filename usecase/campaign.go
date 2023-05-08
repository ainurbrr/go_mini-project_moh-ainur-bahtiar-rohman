package usecase

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/repository/database"

	config "github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/config"
	middlewares "github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/middlewares"
	models "github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/models"
	"github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/models/payload"

	"github.com/gosimple/slug"
	"github.com/labstack/echo/v4"
)

func GetCampaigns(userId int) (interface{}, error) {
	if userId != 0 {
		campaigns, err := database.FindCampaignByUserId(userId)
		if err != nil {
			return campaigns, err
		}
		return campaigns, nil
	}

	campaigns, err := database.FindAllCampaign()
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func GetCampaign(id int) (models.Campaign, error) {

	campaign, err := database.FindCampaignById(id)
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func CreateCampaign(c echo.Context, req *payload.CreateCampaignRequest) (campaign models.Campaign, err error) {

	id, err := middlewares.ExtractTokenId(c)
	if err != nil {
		return campaign, err
	}

	strSlug := fmt.Sprintf("%s %d", req.Name, id)
	time, _ := time.Parse("02/01/2006", c.FormValue("end_date"))
	slug := slug.Make(strSlug)

	campaign = models.Campaign{
		Name:             req.Name,
		UserID:           id,
		ShortDescription: req.ShortDescription,
		Description:      req.Description,
		GoalAmount:       req.GoalAmount,
		Slug:             slug,
		EndDate:          time,
	}

	err = database.CreateCampaign(campaign)
	if err != nil {
		return
	}

	campaignResult, err := database.FindCampaignBySlug(slug)
	if err != nil {
		return
	}

	return campaignResult, nil
}

func UpdateCampaign(c echo.Context) (campaign models.Campaign, err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	campaign, err = database.FindCampaignById(id)
	if err != nil {
		return
	}
	idFromToken, err := middlewares.ExtractTokenId(c)
	if err != nil {
		return campaign, err
	}

	c.Bind(&campaign)

	if campaign.UserID != idFromToken {
		return campaign, errors.New("Unauthorized")
	}

	time, _ := time.Parse("02/01/2006", c.FormValue("end_date"))
	campaign.EndDate = time

	if err = database.UpdateCampaign(&campaign); err != nil {
		return
	}

	return campaign, nil
}

func UploadCampaignImage(c echo.Context) (campaignImages models.Campaign_image, err error) {
	id := c.FormValue("campaign_id")
	campaign_id, _ := strconv.Atoi(id)
	campaign, err := database.FindCampaignById(campaign_id)
	if err != nil {
		return
	}

	c.Bind(&campaign)
	idFromToken, err := middlewares.ExtractTokenId(c)
	if err != nil {
		return campaignImages, err
	}
	if campaign.UserID != idFromToken {
		return campaignImages, errors.New("Unauthorized")
	}

	file, err := c.FormFile("file_name")
	if err != nil {
		return campaignImages, err
	}

	campaignImages = models.Campaign_image{}
	path := fmt.Sprintf("images/campaignImages/%d-%s", campaign_id, file.Filename)
	c.Bind(&campaignImages)
	campaignImages.FileName = path
	if campaignImages.IsPrimary == 1 {
		_, err := MarkAllImagesAsNonPrimary(campaign_id)
		if err != nil {
			return campaignImages, err
		}
	}

	//upload the image
	src, err := file.Open()
	if err != nil {
		return campaignImages, err
	}
	defer src.Close()
	// Create a new file on disk
	dst, err := os.Create(path)
	if err != nil {
		return campaignImages, err
	}
	defer dst.Close()
	// Copy the uploaded file to the destination file
	if _, err = io.Copy(dst, src); err != nil {
		return campaignImages, err
	}

	//save to db
	if err = database.UploadCampaignImage(campaignImages); err != nil {
		return campaignImages, err
	}

	return campaignImages, nil
}

func MarkAllImagesAsNonPrimary(campaignId int) (bool, error) {
	campaign_image := models.Campaign_image{}
	if err := config.DB.Model(&campaign_image).Where("campaign_id = ?", campaignId).Update("is_primary", 0).Error; err != nil {
		return false, err
	}
	return true, nil
}
