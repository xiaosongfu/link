package swagger

// 操作成功的响应对象 写点什么注释好呢
// swagger:response BasicSuccessResponse
type swaggerBasicSuccessResponse struct {
	// in:body
	Body struct {
		// 1 - 操作成功
		Code int `json:"code"`
		// 成功提示
		Message string `json:"message"`
	}
}

// 操作失败的响应对象 写点什么注释好呢
// swagger:response BasicFailedResponse
type swaggerBasicFailedResponse struct {
	// in:body
	Body struct {
		// 0 - 操作失败
		Code int `json:"code"`
		// 失败提示
		Message string `json:"message"`
	}
}
