package swagger

import "xiaosongfu.com/link/dao"

// 添加 category 的请求参数定义 写点什么注释好呢
// swagger:parameters AddCategoryRequest
type swaggerAddCategoryRequest struct {
	// 类别 ID
	//
	// in:query
	CategoryId int `json:"categoryId"`
	// 类别名称
	//
	// in:query
	CategoryName string `json:"categoryName"`
}

// -----------------------------------------------

// 获取所有 category 成功的响应参数定义 写点什么注释好呢
// swagger:response GetCategorysResponse
type swaggerGetCategorysResponse struct {
	// in:body
	Body struct{
		// 1 - 操作成功
		Code int `json:"code"`
		// 成功提示
		Message string `json:"message"`
		// category slice
		Data []dao.Category `json:"data"`
	}
}