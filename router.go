package main

import (
	"net/http"
	"strconv"

	"github.com/xiaosongfu/link/data"
)

// 添加 category
func addCategory(writer http.ResponseWriter, request *http.Request) {
	// 获取请求参数
	var category data.Category
	var err error
	resp := data.BasicResponse{}
	// 解析并判断 categoryId
	category.CategoryId, err = strconv.Atoi(request.PostFormValue("categoryId"))
	if err != nil { // 如果没有传 categoryId 或为 ""，则直接返回并提示错误
		resp.Code = data.ResponseCodeFailed
		resp.Message = "categoryId is not correct"
		WriteJsonResponse(writer, resp)
		return
	}
	// 解析并判断 categoryName
	category.CategoryName = request.PostFormValue("categoryName")
	if StringIsEmpty(category.CategoryName) {
		resp.Code = data.ResponseCodeFailed
		resp.Message = "category is empty"
		WriteJsonResponse(writer, resp)
		return
	}
	// 执行插入
	_, err = category.InsertCategory()
	if err != nil {
		resp.Code = data.ResponseCodeFailed
		resp.Message = err.Error()
	} else {
		resp.Code = data.ResponseCodeSuccess
		resp.Message = "add category success"
	}
	WriteJsonResponse(writer, resp)
}

// 获取所有 category
func getCategorys(writer http.ResponseWriter, request *http.Request) {
	// 执行查询
	categorys, err := data.SelectCategorys()
	categoryListResp := data.CategoryListResponse{}
	if err != nil {
		categoryListResp.Code = data.ResponseCodeFailed
		categoryListResp.Message = err.Error()
	} else {
		categoryListResp.Code = data.ResponseCodeSuccess
		categoryListResp.Message = "get category success"
		categoryListResp.Data = categorys
	}
	WriteJsonResponse(writer, categoryListResp)
}

// 添加 link
func addLink(writer http.ResponseWriter, request *http.Request) {
	// 解析请求参数
	request.ParseForm()
	// 获取请求参数
	var link data.Link
	resp := data.BasicResponse{}
	var err error
	// 解析并判断 url
	link.Url = request.PostForm.Get("url")
	if !IsCorrectUrl(link.Url) {
		resp.Code = data.ResponseCodeFailed
		resp.Message = "url is not correct"
		WriteJsonResponse(writer, resp)
		return
	}
	// 解析并判断 categoryId
	link.CategoryId, err = strconv.Atoi(request.PostForm.Get("categoryId"))
	if err != nil {
		link.CategoryId = data.DefaultCategoryId
	}
	// 获取 tag
	link.Tag = request.PostForm.Get("tag")
	// 获取网页 title
	link.Title, _ = QueryLinkTitle(link.Url)
	// 执行插入
	_, err = link.InsertLink()
	if err != nil {
		resp.Code = data.ResponseCodeFailed
		resp.Message = err.Error()
	} else {
		resp.Code = data.ResponseCodeSuccess
		resp.Message = "add link success"
	}
	WriteJsonResponse(writer, resp)
}

// 获取所有 link
func getLinks(writer http.ResponseWriter, request *http.Request) {
	// 开始查询
	links, err := data.SelectLinks()
	linkListResp := data.LinkListResponse{}
	if err != nil {
		linkListResp.Code = data.ResponseCodeFailed
		linkListResp.Message = err.Error()
	} else {
		linkListResp.Code = data.ResponseCodeSuccess
		linkListResp.Message = "get category success"
		linkListResp.Data = links
	}
	WriteJsonResponse(writer, linkListResp)
}

// 根据 categoryId 获取 link
func getLinksByCategoryId(writer http.ResponseWriter, request *http.Request) {
	// 解析请求参数
	request.ParseForm()
	// 获取请求参数
	var categoryId int
	var err error
	linkListResp := data.LinkListResponse{}
	categoryId, err = strconv.Atoi(request.Form.Get("categoryId"))
	if err != nil {
		linkListResp.Code = data.ResponseCodeFailed
		linkListResp.Message = err.Error()
		WriteJsonResponse(writer, linkListResp)
		return
	}
	// 开始查询
	links, err := data.SelectLinksByCategoryId(categoryId)
	if err != nil {
		linkListResp.Code = data.ResponseCodeFailed
		linkListResp.Message = err.Error()
	} else {
		linkListResp.Code = data.ResponseCodeSuccess
		linkListResp.Message = "get category success"
		linkListResp.Data = links
	}
	WriteJsonResponse(writer, linkListResp)
}
