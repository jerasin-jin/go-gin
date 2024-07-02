package model

type ProductCategoryRequest struct {
	Name        string `json:"name" binding:"required" gorm:"unique"`
	Description string `json:"description"`
}
type ProductCategory struct {
	BaseModel
	Name        string `json:"name" binding:"required" gorm:"unique"`
	Description string `json:"description"`
}
