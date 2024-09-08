package model

type ProductCategory struct {
	BaseModel
	Name        string `json:"name" binding:"required" gorm:"unique"`
	Description string `json:"description"`
	Products    []Product
}

func (ProductCategory) TableName() string {
	return "productCategories"
}
