package request

type UserRequest struct {
	Username string `json:"username" binding:"required" example:"admin"`
	Password string `json:"password" binding:"required" example:"1234"`
	Fullname string `json:"fullname" binding:"required" example:"admin test"`
	Avatar   string `json:"avatar" example:"admin"`
}
