package request

type OrderRequest struct {
	Orders []OrderItem `json:"orders" binding:"required"`
}

type OrderItem struct {
	Id                int     `json:"id" binding:"required" example:"1"`
	Name              string  `json:"name" binding:"required" example:"apple"`
	Description       string  `json:"description" example:"apple"`
	Price             float64 `json:"price" binding:"required" example:"200"`
	Amount            int     `json:"amount" binding:"required" example:"10"`
	ProductCategoryId int     `json:"product_category_id" binding:"required" example:"1"`
}
