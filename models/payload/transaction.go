package payload

import "github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/models"

type CreateTransactionRequest struct {
	CampaignID int    `json:"campaign_id" form:"campaign_id"`
	Amount     int    `json:"amount" form:"amount"`
	Status     string `json:"status" form:"status"`
	Code       string `json:"code" form:"code"`
	PaymentURL string `json:"payment_url"`
	User       models.User
}
