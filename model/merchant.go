package model

type Merchant struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
}

func (Merchant) TableName() string {
	return "merchants"
}