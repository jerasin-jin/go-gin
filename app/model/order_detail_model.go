package model

type OrderDetail struct {
	BaseModel
	ProductID         uint    `json:"id" binding:"required"`
	Name              string  `json:"name" binding:"required"`
	Description       string  `json:"description"`
	Price             float64 `json:"price" binding:"required"`
	Amount            int     `json:"amount" binding:"required"`
	ProductCategoryID uint    `json:"product_category_id" binding:"required"`
	OrderID           uint    `json:"order_id" binding:"required"`
}

func (OrderDetail) TableName() string {
	return "orderDetails"
}
