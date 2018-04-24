package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/xiaosongfu/link/data"
)

// 添加 category
func addCategory(writer http.ResponseWriter, request *http.Request) {
	// 获取请求参数
	var category data.Category
	var err error
	category.CategoryId, err = strconv.Atoi(request.PostFormValue("categoryId"))
	if err != nil { // 如果没有传 categoryId 或为 ""，则直接返回并提示错误
		writer.Write([]byte(err.Error()))
		return
	}
	category.CategoryName = request.PostFormValue("categoryName")
	// 执行插入
	_, err = category.InsertCategory()
	resp := data.BasicResponse{}
	if err != nil {
		resp.Code = data.ResponseCodeFailed
		resp.Message = err.Error()
	} else {
		resp.Code = data.ResponseCodeSuccess
		resp.Message = "add category success."
	}
	writer.Header().Set("Content-Type", "application/json")
	jsonResp, _ := json.Marshal(resp)
	writer.Write(jsonResp)
}

// 获取所有 category
func getCategorys(writer http.ResponseWriter, request *http.Request) {
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
	writer.Header().Set("Content-Type", "application/json")
	jsonResp, _ := json.Marshal(categoryListResp)
	writer.Write(jsonResp)
}

// 添加 link
func addLink(writer http.ResponseWriter, request *http.Request) {
	// 解析请求参数
	request.ParseForm()
	// 获取请求参数
	var link data.Link
	var err error
	link.CategoryId, err = strconv.Atoi(request.PostForm.Get("categoryId"))
	if err != nil {
		link.CategoryId = data.DefaultCategoryId
	}
	link.Url = request.PostForm.Get("url")
	link.Tag = request.PostForm.Get("tag")
	// 执行插入
	_, err = link.InsertLink()
	resp := data.BasicResponse{}
	if err != nil {
		resp.Code = data.ResponseCodeFailed
		resp.Message = err.Error()
	} else {
		resp.Code = data.ResponseCodeSuccess
		resp.Message = "add link success"
	}
	writer.Header().Set("Content-Type", "application/json")
	jsonResp, _ := json.Marshal(resp)
	writer.Write(jsonResp)
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
	writer.Header().Set("Content-Type", "application/json")
	jsonResp, _ := json.Marshal(linkListResp)
	writer.Write(jsonResp)
}

// 根据 categoryId 获取 link
func getLinksByCategoryId(writer http.ResponseWriter, request *http.Request) {
	// 解析请求参数
	request.ParseForm()
	// 获取请求参数
	var categoryId int
	var err error
	categoryId, err = strconv.Atoi(request.Form.Get("categoryId"))
	if err != nil {
		categoryId = data.DefaultCategoryId
	}
	// 开始查询
	links, err := data.SelectLinksByCategoryId(categoryId)
	linkListResp := data.LinkListResponse{}
	if err != nil {
		linkListResp.Code = data.ResponseCodeFailed
		linkListResp.Message = err.Error()
	} else {
		linkListResp.Code = data.ResponseCodeSuccess
		linkListResp.Message = "get category success"
		linkListResp.Data = links
	}
	writer.Header().Set("Content-Type", "application/json")
	jsonResp, _ := json.Marshal(linkListResp)
	writer.Write(jsonResp)
}
