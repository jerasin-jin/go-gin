package model

type OrderDetail struct {
	BaseModel
	ProductID uint    `json:"product_id" binding:"required"`
	Price     float64 `json:"price" binding:"required"`
	Amount    int     `json:"amount" binding:"required"`
	OrderID   uint    `json:"order_id" binding:"required"`
}

func (OrderDetail) TableName() string {
	return "orderDetails"
}
