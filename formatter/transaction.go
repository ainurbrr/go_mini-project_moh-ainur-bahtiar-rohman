package formatter

import (
	"penggalangan-dana/models"
	"time"
)

type CampaignTransactionFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatCampaignTransaction(transaction models.Transaction) CampaignTransactionFormatter{
	formatter := CampaignTransactionFormatter{}

	formatter.ID = transaction.ID
	formatter.Name = transaction.User.Name
	formatter.Amount = transaction.Amount
	formatter.CreatedAt = transaction.CreatedAt
	return formatter
}

func FormatCampaignTransactions(transactions []models.Transaction) []CampaignTransactionFormatter{
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
