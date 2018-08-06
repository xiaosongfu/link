package swagger

import "xiaosongfu.com/link/dao"

// 添加 link 的请求参数定义 写点什么注释好呢
// swagger:parameters AddLinkRequest
type swaggerAddLinkRequest struct {
	// 网页链接
	//
	// in:query
	Url string `json:"url"`
	// 类别 ID
	//
	// in:query
	CategoryId string `json:"categoryId"`
	// 标签
	//
	// in:query
	Tag string `json:"tag"`
}

// 根据 categoryId 获取 link 的请求参数定义 写点什么注释好呢
// swagger:parameters GetLinksByCategoryIdRequest
type swaggerGetLinksByCategoryIdRequest struct {
	// 类别 ID
	//
	// in:query
	CategoryId string `json:"categoryId"`
}

// -----------------------------------------------

// 获取 link 成功的响应参数定义 写点什么注释好呢
// swagger:response GetLinksResponse
type swaggerGetLinksResponse struct {
	// in:body
	Body struct {
		// 1 - 操作成功
		Code int `json:"code"`
		// 成功提示
		Message string `json:"message"`
		// link slice
		Data []dao.Link `json:"data"`
	}
}
