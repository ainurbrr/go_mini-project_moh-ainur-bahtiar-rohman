package models

import (
	"time"

	"gorm.io/gorm"
)

type Campaign struct {
	gorm.Model
	ID               int       `json:"id" form:"id"`
	UserID           int       `json:"user_id" form:"user_id"`
	Name             string    `json:"name" form:"name" binding:"required"`
	ShortDescription string    `json:"short_description" form:"short_description" binding:"required"`
	Description      string    `json:"description" form:"description" binding:"required"`
	BackerCount      int       `json:"backer_count" form:"backer_count"`
	GoalAmount       int       `json:"goal_amount" form:"goal_amount" binding:"required"`
	TotalAmount      int       `json:"total_amount" form:"total_amount"`
	Slug             string    `json:"slug" form:"slug"`
	EndDate          time.Time `json:"end_date" form:"end_date" binding:"required"`
	User             *User
	Transactions     []*Transaction
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
