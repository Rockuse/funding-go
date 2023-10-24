package transaction

type TransactionFormat struct {
	Id           int    `json:"id"`
	Code         string `json:"code"`
	CampaignName string `json:"campaign_name"`
	UserId       int    `json:"user_id"`
	Amount       int    `json:"amount"`
	Status       string `json:"status"`
}

func FormatTransaction(data Transaction) TransactionFormat {
	var format TransactionFormat
	format.Id = data.Id
	format.Code = data.Code
	format.CampaignName = data.Campaign.Name
	format.Amount = data.Amount
	format.Status = data.Status
	return format
}

func FormatListTransaction(data []Transaction) []TransactionFormat {
	var format []TransactionFormat
	for _, trans := range data {
		format = append(format, FormatTransaction(trans))
	}

	return format
}
