package request

type UserRequest struct {
	Username string `json:"username" binding:"required" example:"admin"`
	Password string `json:"password" binding:"required" example:"1234"`
	Fullname string `json:"fullname" binding:"required" example:"admin test"`
	Avatar   string `json:"avatar" example:"admin"`
}

type UpdateUserRequest struct {
	Username string `json:"username" example:"admin"`
	// Password string `json:"password" example:"1234"`
	Fullname string `json:"fullname" example:"admin test"`
	Avatar   string `json:"avatar" example:"admin"`
}
