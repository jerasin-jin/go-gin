package model

type ProductCategory struct {
	BaseModel
	Name        string `gorm:"unique"`
	Description string
	Products    []Product
}

func (ProductCategory) TableName() string {
	return "productCategories"
}
