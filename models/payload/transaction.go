package payload

import "struktur-penggalangan-dana/models"

type CreateTransactionRequest struct {
	CampaignID int    `json:"campaign_id" form:"campaign_id"`
	Amount     int    `json:"amount" form:"amount"`
	Status     string `json:"status" form:"status"`
	Code       string `json:"code" form:"code"`
	PaymentURL string `json:"payment_url"`
	User       models.User
}
