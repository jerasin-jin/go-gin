package request

type ProductCategoryRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type UpdateProductCategory struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
