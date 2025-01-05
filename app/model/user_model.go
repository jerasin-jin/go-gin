package model

type User struct {
	BaseModel
	Username string
	Password string
	Fullname string
	Avatar   string
	Order    []Order `gorm:"foreignKey:CreatedBy;references:ID"`
}
