package model

type Order struct {
	BaseModel
	TotalPrice  float64
	TotalAmount int
	OrderDetail []OrderDetail
	CreatedBy   uint
}
