package response

type PaginationResponse struct {
	ResponseKey     string `json:"response_key"`
	ResponseMessage string `json:"response_message"`
	TotalPage       int    `json:"totalPage"`
	Page            int    `json:"page"`
	PageSize        int    `json:"pageSize"`
}
