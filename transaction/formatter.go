package transaction

import "time"

type CampaignTransactionFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatCampaignTransaction(transaction Transaction) CampaignTransactionFormatter {
	formatter := CampaignTransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.Name = transaction.User.Name
	formatter.Amount = transaction.Amount
	formatter.CreatedAt = transaction.CreatedAt
	return formatter
}

func FormatCampaignTransactions(transactions []Transaction) []CampaignTransactionFormatter {
	if len(transactions) == 0 {
		return []CampaignTransactionFormatter{}
	}

	var transactionsFormatter []CampaignTransactionFormatter

	for _, transaction := range transactions {
		formatter := FormatCampaignTransaction(transaction)
		transactionsFormatter = append(transactionsFormatter, formatter)
	}

	return transactionsFormatter
}

type UserTransactionFromatter struct {
	ID        int              `json:"id"`
	Amount    int              `json:"amount"`
	Status    string           `json:"status"`
	CreatedAt time.Time        `json:"created_at"`
	Campaign  CampaignFormater `json:"campaign"`
}

type CampaignFormater struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func FormatUserTransaction(transaction Transaction) UserTransactionFromatter {
	formatter := UserTransactionFromatter{}
	formatter.ID = transaction.ID
	formatter.Amount = transaction.Amount
	formatter.Status = transaction.Stauts
	formatter.CreatedAt = transaction.CreatedAt

	campaignFormater := CampaignFormater{}
	campaignFormater.Name = transaction.Campaign.Name
	campaignFormater.ImageURL = ""
	if len(transaction.Campaign.CampaignImages) > 0 {
		campaignFormater.ImageURL = transaction.Campaign.CampaignImages[0].FileName
	}

	formatter.Campaign = campaignFormater

	return formatter
}

func FormatUserTransactions(transactions []Transaction) []UserTransactionFromatter {
	if len(transactions) == 0 {
		return []UserTransactionFromatter{}
	}

	var transactionsFormatter []UserTransactionFromatter

	for _, transaction := range transactions {
		formatter := FormatUserTransaction(transaction)
		transactionsFormatter = append(transactionsFormatter, formatter)
	}

	return transactionsFormatter
}
