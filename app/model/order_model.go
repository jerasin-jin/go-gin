package model

type Order struct {
	BaseModel
	TotalPrice  float64 `gorm:"not null"`
	TotalAmount int     `gorm:"not null"`
	OrderDetail []OrderDetail
	CreatedBy   uint `gorm:"not null"`
}
