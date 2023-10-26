package model

type Customer struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	RoleID   string `json:"role_id"`
}

func (Customer) TableName() string {
	return "customers"
}
