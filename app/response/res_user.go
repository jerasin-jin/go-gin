package response

type User struct {
	Username string `json:"username"`
	// Password string `json:"password"`
	FullName string `json:"fullName"`
	Avatar   string `json:"avatar"`
}
