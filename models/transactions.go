package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ID         int    `json:"id" form:"id"`
	CampaignID int    `json:"campaign_id" form:"campaign_id"`
	UserID     int    `json:"user_id" form:"user_id"`
	Amount     int    `json:"amount" form:"amount"`
	Status     string `json:"status" form:"status"`
	Code       string `json:"code" form:"code"`
	PaymentURL string `json:"payment_url" form:"payment_url"`
	User       *User
	Campaign   *Campaign
}
