package api

import (
	"net/http"
	"strconv"

	"ffll.fun/link/dao"
	"ffll.fun/link/handlers"
	"ffll.fun/link/service"
	"ffll.fun/link/util"
)

// link 列表响应
type LinkListResponse struct {
	handlers.BasicResponse
	Data []dao.Link `json:"data"`
}

// 添加 link
//
// swagger:route POST /api/v1/addLink link AddLinkRequest
// 该接口实现添加 link 功能.
// 该填点什么描述好呢.
// responses:
//   0: BasicFailedResponse
//   1: BasicSuccessResponse
func AddLink(writer http.ResponseWriter, request *http.Request) {
	// 解析请求参数
	request.ParseForm()
	// 获取请求参数
	var link dao.Link
	resp := handlers.BasicResponse{}
	var err error
	// 解析并判断 url
	link.Url = request.PostForm.Get("url")
	if !util.IsCorrectUrl(link.Url) {
		resp.Code = handlers.ResponseCodeFailed
		resp.Message = "url is not correct"
		util.WriteJsonResponse(writer, resp)
		return
	}
	// 解析并判断 categoryId
	link.CategoryId, err = strconv.Atoi(request.PostForm.Get("categoryId"))
	if err != nil {
		link.CategoryId = dao.DefaultCategoryId
	}
	// 获取 tag
	link.Tag = request.PostForm.Get("tag")
	// 获取网页 title
	link.Title, _ = util.QueryLinkTitle(link.Url)
	// 执行插入
	_, err = service.AddLink(&link)
	if err != nil {
		resp.Code = handlers.ResponseCodeFailed
		resp.Message = err.Error()
	} else {
		resp.Code = handlers.ResponseCodeSuccess
		resp.Message = "add link success.title: " + link.Title
	}
	util.WriteJsonResponse(writer, resp)
}

// 获取所有 link
//
// swagger:route GET /api/v1/getLinks link GetLinks
// 该接口实现了查询所有 link 功能.
// 该填点什么描述好呢.
// responses:
//   0: BasicFailedResponse
//   1: GetLinksResponse
func GetLinks(writer http.ResponseWriter, request *http.Request) {
	// 开始查询
	links, err := service.GetLinks()
	linkListResp := LinkListResponse{}
	if err != nil {
		linkListResp.Code = handlers.ResponseCodeFailed
		linkListResp.Message = err.Error()
	} else {
		linkListResp.Code = handlers.ResponseCodeSuccess
		linkListResp.Message = "get link success"
		linkListResp.Data = links
	}
	util.WriteJsonResponse(writer, linkListResp)
}

// 根据 categoryId 获取 link
//
// swagger:operation GET /api/v1/getLinksByCategoryId link GetLinksByCategoryId
// ---
// summary: 根据 categoryId 获取 link.
// description: 该填点什么描述好呢.
// parameters:
// - name: categoryId
//   in: query
//   description: 类别 ID
//   type: string
//   required: true
// responses:
//   "0":
//     "$ref": "#/responses/BasicFailedResponse"
//   "1":
//     "$ref": "#/responses/GetLinksResponse"
func GetLinksByCategoryId(writer http.ResponseWriter, request *http.Request) {
	// 解析请求参数
	request.ParseForm()
	// 获取请求参数
	var categoryId int
	var err error
	linkListResp := LinkListResponse{}
	categoryId, err = strconv.Atoi(request.Form.Get("categoryId"))
	if err != nil {
		linkListResp.Code = handlers.ResponseCodeFailed
		linkListResp.Message = err.Error()
		util.WriteJsonResponse(writer, linkListResp)
		return
	}
	// 开始查询
	links, err := service.GetLinksByCategoryId(categoryId)
	if err != nil {
		linkListResp.Code = handlers.ResponseCodeFailed
		linkListResp.Message = err.Error()
	} else {
		linkListResp.Code = handlers.ResponseCodeSuccess
		linkListResp.Message = "get link success"
		linkListResp.Data = links
	}
	util.WriteJsonResponse(writer, linkListResp)
}
