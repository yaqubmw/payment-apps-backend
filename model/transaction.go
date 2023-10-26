package model

type Transaction struct {
	ID         string    `json:"id"`
	CustomerID string    `json:"customer_id"`
	MerchantID string    `json:"merchant_id"`
	Amount     float64   `json:"amount"`
}

func (Transaction) TableName() string {
	return "tx_transaction"
}
