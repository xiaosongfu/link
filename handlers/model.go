package handlers

// 失败响应 code
const ResponseCodeFailed = 0

// 成功响应 code
const ResponseCodeSuccess = 1

// 基础响应
type BasicResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
