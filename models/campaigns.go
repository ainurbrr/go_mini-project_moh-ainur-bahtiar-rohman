package models

import (
	"gorm.io/gorm"
)

type Campaign struct {
	gorm.Model
	ID               int    `json:"id" form:"id"`
	UserID           int    `json:"user_id" form:"user_id"`
	Name             string `json:"name" form:"name"`
	ShortDescription string `json:"short_description" form:"short_description"`
	Description      string `json:"description" form:"description"`
	BackerCount      int    `json:"backer_count" form:"backer_count"`
	GoalAmount       int    `json:"goal_amount" form:"goal_amount"`
	TotalAmount      int    `json:"total_amount" form:"total_amount"`
	User             *User
	Transactions      []*Transaction
	CampaignImages   []*Campaign_image
}

type Campaign_image struct {
	gorm.Model
	ID         int    `json:"id" form:"id"`
	CampaignID int    `json:"campaign_id" form:"campaign_id"`
	FileName   string `json:"file_name" form:"file_name"`
	IsPrimary  int    `json:"is_primary" form:"is_primary"`
	Campaign   *Campaign
}
