package request

type RoleInfoRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type UpdateRoleInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
