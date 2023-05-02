package formatter

import (
	"penggalangan-dana/models"
	"time"
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
