package formatter

import (
	"time"

	"github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/models"
)

type CampaignTransactionFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatCampaignTransaction(transaction models.Transaction) CampaignTransactionFormatter {
	formatter := CampaignTransactionFormatter{}

	formatter.ID = transaction.ID
	formatter.Name = transaction.User.Name
	formatter.Amount = transaction.Amount
	formatter.CreatedAt = transaction.CreatedAt
	return formatter
}

func FormatCampaignTransactions(transactions []models.Transaction) []CampaignTransactionFormatter {
	if len(transactions) == 0 {
		return []CampaignTransactionFormatter{}
	}
	formatters := []CampaignTransactionFormatter{}

	for _, transaction := range transactions {
		formatter := FormatCampaignTransaction(transaction)
		formatters = append(formatters, formatter)
	}
	return formatters
}

type UserTransactionFormatter struct {
	ID        int                               `json:"id"`
	Amount    int                               `json:"amount"`
	Status    string                            `json:"status"`
	CreatedAt time.Time                         `json:"created_at"`
	Campaign  UserTransactionsCampaignFormatter `json:"campaign"`
}
type UserTransactionsCampaignFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func FormatUserTransaction(transaction models.Transaction) UserTransactionFormatter {
	formatter := UserTransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.Amount = transaction.Amount
	formatter.Status = transaction.Status
	formatter.CreatedAt = transaction.CreatedAt

	campaignFormatter := UserTransactionsCampaignFormatter{}
	campaignFormatter.Name = transaction.Campaign.Name
	campaignFormatter.ImageURL = ""

	if len(transaction.Campaign.CampaignImages) > 0 {
		campaignFormatter.ImageURL = transaction.Campaign.CampaignImages[0].FileName
	}
	formatter.Campaign = campaignFormatter
	return formatter
}

func FormatUserTransactions(transactions []models.Transaction) []UserTransactionFormatter {
	if len(transactions) == 0 {
		return []UserTransactionFormatter{}
	}
	formatters := []UserTransactionFormatter{}

	for _, transaction := range transactions {
		formatter := FormatUserTransaction(transaction)
		formatters = append(formatters, formatter)
	}
	return formatters
}

type TransactionFormatter struct {
	ID         int       `json:"id"`
	CampaignID int       `json:"campaign_id"`
	UserID     int       `json:"user_id"`
	Amount     int       `json:"amount"`
	Code       string    `json:"code"`
	Status     string    `json:"status"`
	PaymentURL string    `json:"payment_url"`
	CreatedAt  time.Time `json:"created_at"`
}

func FormatTransaction(transaction models.Transaction) TransactionFormatter {
	formatter := TransactionFormatter{}

	formatter.ID = transaction.ID
	formatter.CampaignID = transaction.CampaignID
	formatter.UserID = transaction.UserID
	formatter.Status = transaction.Status
	formatter.Amount = transaction.Amount
	formatter.Code = transaction.Code
	formatter.PaymentURL = transaction.PaymentURL

	return formatter
}
