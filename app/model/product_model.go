package model

import "time"

type Product struct {
	BaseModel
	Name              string `gorm:"unique" json:"name" binding:"required"`
	Description       string
	Price             float64
	Amount            int
	SaleOpenDate      time.Time
	SaleCloseDate     time.Time
	ProductCategoryID uint
	ImgUrl            string `gorm:"default:null"`
}
