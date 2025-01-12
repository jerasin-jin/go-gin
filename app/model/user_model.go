package model

type User struct {
	BaseModel
	Username   string `gorm:"not null"`
	Password   string `gorm:"not null"`
	Fullname   string `gorm:"not null"`
	Avatar     string
	Email      string  `gorm:"not null"`
	Order      []Order `gorm:"foreignKey:CreatedBy;references:ID"`
	RoleInfoID uint    `gorm:"not null"`
}
