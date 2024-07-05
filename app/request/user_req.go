package request

type UserRequest struct {
	Username string `json:"username" binding:"required" gorm:"unique"`
	Password string `json:"password" binding:"required" gorm:"unique"`
	Fullname string `json:"fullname" binding:"required" gorm:"unique"`
	Avatar   string `json:"avatar" binding:"required" gorm:"unique"`
}
