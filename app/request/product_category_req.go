package request

type ProductCategoryRequest struct {
	Name        string `json:"name" binding:"required" gorm:"unique"`
	Description string `json:"description"`
}
