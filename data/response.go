package data

// 失败响应 code
const ResponseCodeFailed = 0

// 成功响应 code
const ResponseCodeSuccess = 1

// 基础响应
type BasicResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// link 列表响应
type LinkListResponse struct {
	BasicResponse
	Data []Link `json:"data"`
}

// category 列表响应
type CategoryListResponse struct {
	BasicResponse
	Data []Category `json:"data"`
}
