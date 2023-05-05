package database

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"penggalangan-dana/config"
	"penggalangan-dana/middlewares"
	"penggalangan-dana/models"
	"strconv"
	"time"

	"github.com/gosimple/slug"
	"github.com/labstack/echo/v4"
)

func FindAll() (interface{}, error) {
	var campaigns []models.Campaign

	if err := config.DB.Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error; err != nil {
		return campaigns, err
	}
	return campaigns, nil

}

func FindByUserId(userId int) (interface{}, error) {
	var campaigns []models.Campaign

	if err := config.DB.Where("user_id = ?", userId).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error; err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func GetCampaigns(userId int) (interface{}, error) {
	if userId != 0 {
		campaigns, err := FindByUserId(userId)
		if err != nil {
			return campaigns, err
		}
		return campaigns, nil
	}

	campaigns, err := FindAll()
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func FindById(id int) (interface{}, error) {
	var campaign models.Campaign

	if err := config.DB.Preload("User").Preload("CampaignImages").Where("id = ?", id).Find(&campaign).Error; err != nil {
		return nil, err
	}
	return campaign, nil
}

func CreateCampaign(c echo.Context) (interface{}, error) {
	campaign := models.Campaign{}
	c.Bind(&campaign)

	id, err := middlewares.ExtractTokenId(c)
	if err != nil {
		return nil, err
	}
	campaign.UserID = id
	strSlug := fmt.Sprintf("%s %d", campaign.Name, campaign.UserID)
	campaign.Slug = slug.Make(strSlug)
	time, _ := time.Parse("02/01/2006", c.FormValue("end_date"))
	campaign.EndDate = time

	if err := config.DB.Create(&campaign).Error; err != nil {
		return nil, err
	}
	return campaign, nil

}

func UpdateCampaign(c echo.Context) (interface{}, error) {
	id, _ := strconv.Atoi(c.Param("id"))
	campaign, err := FindById(id)
	if err != nil {
		return nil, err
	}
	idFromToken, err := middlewares.ExtractTokenId(c)
	if err != nil {
		return nil, err
	}

	campaignModel := campaign.(models.Campaign)

	c.Bind(&campaignModel)

	if campaignModel.UserID != idFromToken {
		return nil, errors.New("Unauthorized")
	}

	time, _ := time.Parse("02/01/2006", c.FormValue("end_date"))
	campaignModel.EndDate = time

	if err := config.DB.Model(&campaignModel).Updates(campaignModel).Error; err != nil {
		return nil, err
	}

	return campaignModel, nil
}

func UploadImage(c echo.Context) (interface{}, error) {
	id := c.FormValue("campaign_id")
	campaign_id, _ := strconv.Atoi(id)
	campaign, err := FindById(campaign_id)
	if err != nil {
		return nil, err
	}
	campaignModel := campaign.(models.Campaign)

	c.Bind(&campaignModel)
	idFromToken, err := middlewares.ExtractTokenId(c)
	if err != nil {
		return nil, err
	}
	if campaignModel.UserID != idFromToken {
		return nil, errors.New("Unauthorized")
	}

	campaignImageModel := models.Campaign_image{}

	file, err := c.FormFile("file_name")
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("images/campaignImages/%d-%s", campaign_id, file.Filename)
	c.Bind(&campaignImageModel)
	campaignImageModel.FileName = path
	if campaignImageModel.IsPrimary == 1 {
		_, err := MarkAllImagesAsNonPrimary(campaign_id)
		if err != nil {
			return nil, err
		}
	}

	//upload the image
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()
	// Create a new file on disk
	dst, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	defer dst.Close()
	// Copy the uploaded file to the destination file
	if _, err = io.Copy(dst, src); err != nil {
		return nil, err
	}

	//save to db
	if err := config.DB.Create(&campaignImageModel).Error; err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return campaignImageModel, nil
}

func MarkAllImagesAsNonPrimary(campaignId int) (bool, error) {
	campaign_image := models.Campaign_image{}
	if err := config.DB.Model(&campaign_image).Where("campaign_id = ?", campaignId).Update("is_primary", 0).Error; err != nil {
		return false, err
	}
	return true, nil
}
