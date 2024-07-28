package model

import "time"

type Product struct {
	BaseModel
	Name          string    `json:"name" binding:"required" gorm:"unique"`
	Description   string    `json:"description"`
	Price         float64   `json:"price" binding:"required"`
	Amount        int       `json:"amount" binding:"required"`
	SaleOpenDate  time.Time `json:"sale_open_date" binding:"required"`
	SaleCloseDate time.Time `json:"sale_close_date" binding:"required"`
}
