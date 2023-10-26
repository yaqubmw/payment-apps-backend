package model

import "time"

type History struct {
	ID            string    `json:"id"`
	TransactionID string    `json:"transactionid"`
	CustomerName  string    `json:"customername"`
	MerchantName  string    `json:"merchantname"`
	Amount        float64   `json:"amount"`
	Timestamp     time.Time `json:"timestamp"`
}
