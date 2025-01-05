package model

type OrderDetail struct {
	BaseModel
	ProductID uint
	Price     float64
	Amount    int
	OrderID   uint
}

func (OrderDetail) TableName() string {
	return "orderDetails"
}
