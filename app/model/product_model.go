package model

import "time"

type Product struct {
	BaseModel
	Name              string    `gorm:"unique" json:"name" binding:"required"`
	Description       string    `json:"description"`
	Price             float64   `json:"price" binding:"required"`
	Amount            int       `json:"amount" binding:"required"`
	SaleOpenDate      time.Time `json:"sale_open_date" binding:"required"`
	SaleCloseDate     time.Time `json:"sale_close_date" binding:"required"`
	ProductCategoryID uint      `json:"product_category_id" binding:"required"`
	ImgUrl            string    `json:"img_url"`
}
