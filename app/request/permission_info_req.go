package request

type PermissionInfoRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type UpdatePermissionInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
