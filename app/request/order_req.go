package request

import "time"

type OrderRequest struct {
	Orders []OrderItem `json:"orders" binding:"required"`
}

type OrderItem struct {
	ProductId int       `json:"product_id" binding:"required" example:"1"`
	Amount    int       `json:"amount" binding:"required" example:"10"`
	UpdatedAt time.Time `json:"updated_at"`
}
