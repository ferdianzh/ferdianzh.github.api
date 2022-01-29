package models

type Work struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	IDRole      string `json:"id_role"`
	Name        string `json:"name"`
	Tech        string `json:"tech"`
	Description string `json:"description"`
}