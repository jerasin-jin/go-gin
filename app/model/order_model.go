package model

type Order struct {
	BaseModel
	TotalPrice  int `json:"total_price" binding:"required"`
	TotalAmount int `json:"total_amount" binding:"required"`
	OrderDetail []OrderDetail
}
