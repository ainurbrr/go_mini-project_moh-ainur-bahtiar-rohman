package formatter

import (
	"errors"
	"time"

	"github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/models"
)

type CampaignFormatter struct {
	ID               int       `json:"id"`
	UserID           int       `json:"user_id"`
	Name             string    `json:"name"`
	ShortDescription string    `json:"short_description"`
	ImageURL         string    `json:"image_url"`
	GoalAmount       int       `json:"goal_amount"`
	TotalAmount      int       `json:"total_amount"`
	EndDate          time.Time `json:"end_date"`
	Slug             string    `json:"slug"`
}

func FormatCampaign(campaign models.Campaign) CampaignFormatter {
	formatter := CampaignFormatter{}
	formatter.ID = campaign.ID
	formatter.UserID = campaign.UserID
	formatter.Name = campaign.Name
	formatter.ShortDescription = campaign.ShortDescription
	formatter.GoalAmount = campaign.GoalAmount
	formatter.TotalAmount = campaign.TotalAmount
	formatter.EndDate = campaign.EndDate
	formatter.Slug = campaign.Slug
	formatter.ImageURL = ""

	if len(campaign.CampaignImages) > 0 {
		formatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	return formatter
}

func FormatCampaigns(campaigns []models.Campaign) []CampaignFormatter {
	campaignsFormatter := []CampaignFormatter{}

	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}
	return campaignsFormatter
}

type CampaignDetailFormatter struct {
	ID               int                      `json:"id"`
	Name             string                   `json:"name"`
	ShortDescription string                   `json:"short_description"`
	Description      string                   `json:"description"`
	ImageURL         string                   `json:"image_url"`
	GoalAmount       int                      `json:"goal_amount"`
	TotalAmount      int                      `json:"total_amount"`
	BackerCount      int                      `json:"backer_count"`
	UserID           int                      `json:"user_id"`
	EndDate          time.Time                `json:"end_date"`
	CloseDate        int                      `json:"close_date"`
	Slug             string                   `json:"slug"`
	User             CampaignUserFormatter    `json:"user"`
	Images           []CampaignImageFormatter `json:"images"`
}

type CampaignUserFormatter struct {
	Name      string `json:"name"`
	Image_URL string `json:"Image_URL"`
}

type CampaignImageFormatter struct {
	Image_URL string `json:"Image_URL"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatCampaignDetail(campaign models.Campaign) (interface{}, error) {
	formatterDetail := CampaignDetailFormatter{}
	formatterDetail.ID = campaign.ID
	formatterDetail.Name = campaign.Name
	formatterDetail.ShortDescription = campaign.ShortDescription
	formatterDetail.Description = campaign.Description
	formatterDetail.GoalAmount = campaign.GoalAmount
	formatterDetail.BackerCount = campaign.BackerCount
	formatterDetail.TotalAmount = campaign.TotalAmount
	formatterDetail.UserID = campaign.UserID
	formatterDetail.EndDate = campaign.EndDate
	ClosingDate := campaign.EndDate.Sub(time.Now())
	formatterDetail.CloseDate = int(ClosingDate.Hours() / 24)
	formatterDetail.Slug = campaign.Slug
	formatterDetail.ImageURL = ""

	if len(campaign.CampaignImages) > 0 {
		formatterDetail.ImageURL = campaign.CampaignImages[0].FileName
	}
	if campaign.ID == 0 {
		return nil, errors.New("campaign not found")
	}

	user := campaign.User
	formatterUserCampaign := CampaignUserFormatter{}
	formatterUserCampaign.Name = user.Name
	formatterUserCampaign.Image_URL = user.Avatar_File_Name
	formatterDetail.User = formatterUserCampaign

	images := []CampaignImageFormatter{}

	for _, image := range campaign.CampaignImages {
		formatterImageCampaign := CampaignImageFormatter{}
		formatterImageCampaign.Image_URL = image.FileName

		isPrimary := false
		if image.IsPrimary == 1 {
			isPrimary = true
		}
		formatterImageCampaign.IsPrimary = isPrimary

		images = append(images, formatterImageCampaign)
	}

	formatterDetail.Images = images
	return formatterDetail, nil
}
