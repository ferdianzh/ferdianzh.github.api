package models

type Works struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	Role        string `json:"role"`
	Tech        string `json:"tech"`
	Description string `json:"description"`
}