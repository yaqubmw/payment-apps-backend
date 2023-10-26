package model

type Role struct {
	ID   string `json:"id"`
	Role string `json:"role"`
}

func (Role) TableName() string {
	return "role"
}