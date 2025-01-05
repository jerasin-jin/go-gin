package request

type OrderRequest struct {
	Orders []OrderItem `json:"orders" binding:"required"`
}

type OrderItem struct {
	ProductId int `json:"product_id" binding:"required" example:"1"`
	Amount    int `json:"amount" binding:"required" example:"10"`
}
