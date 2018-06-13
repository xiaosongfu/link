package api

import (
	"net/http"
	"strconv"

	"ffll.fun/link/dao"
	"ffll.fun/link/handlers"
	"ffll.fun/link/service"
	"ffll.fun/link/util"
)

// category 列表响应
type CategoryListResponse struct {
	handlers.BasicResponse
	Data []dao.Category `json:"data"`
}

// 添加 category
//
// swagger:route POST /api/v1/addCategory category AddCategoryRequest
// 该接口实现添加 category 功能.
// 该填点什么描述好呢.
// responses:
//   0: BasicFailedResponse
//   1: BasicSuccessResponse
func AddCategory(writer http.ResponseWriter, request *http.Request) {
	// 获取请求参数
	var category dao.Category
	var err error
	resp := handlers.BasicResponse{}
	// 解析并判断 categoryId
	category.CategoryId, err = strconv.Atoi(request.PostFormValue("categoryId"))
	if err != nil { // 如果没有传 categoryId 或为 ""，则直接返回并提示错误
		resp.Code = handlers.ResponseCodeFailed
		resp.Message = "category id is not correct"
		util.WriteJsonResponse(writer, resp)
		return
	}
	// 解析并判断 categoryName
	category.CategoryName = request.PostFormValue("categoryName")
	if util.StringIsEmpty(category.CategoryName) {
		resp.Code = handlers.ResponseCodeFailed
		resp.Message = "category name is empty"
		util.WriteJsonResponse(writer, resp)
		return
	}
	// 执行插入
	_, err = service.AddCategory(&category)
	if err != nil {
		resp.Code = handlers.ResponseCodeFailed
		resp.Message = err.Error()
	} else {
		resp.Code = handlers.ResponseCodeSuccess
		resp.Message = "add category success"
	}
	util.WriteJsonResponse(writer, resp)
}

// 获取所有 category
//
// swagger:route GET /api/v1/getCategorys category GetCategorys
// 该接口实现获取所有 category 功能.
// 该填点什么描述好呢.
// responses:
//   0: BasicFailedResponse
//   1: GetCategorysResponse
func GetCategorys(writer http.ResponseWriter, request *http.Request) {
	// 执行查询
	categorys, err := service.GetCategorys()
	categoryListResp := CategoryListResponse{}
	if err != nil {
		categoryListResp.Code = handlers.ResponseCodeFailed
		categoryListResp.Message = err.Error()
	} else {
		categoryListResp.Code = handlers.ResponseCodeSuccess
		categoryListResp.Message = "get category success"
		categoryListResp.Data = categorys
	}
	util.WriteJsonResponse(writer, categoryListResp)
}
