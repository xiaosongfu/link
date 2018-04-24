package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/xiaosongfu/link/data"
	"log"
	"net/http"
	"strconv"
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
	insertId, err := category.InsertCategory()
	if err != nil {
		writer.Write([]byte(err.Error()))
	} else {
		log.Println("insertId: " + string(insertId))
		writer.Write([]byte("ok"))
	}
}

// 获取所有 category
func getCategorys(writer http.ResponseWriter, request *http.Request) {
	categorys, err := data.SelectCategorys()
	if err != nil {
		writer.Write([]byte(err.Error()))
	} else {
		spew.Dump(categorys)
		writer.Write([]byte("ok"))
	}
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
	insertId, err := link.InsertLink()
	if err != nil {
		writer.Write([]byte(err.Error()))
	} else {
		log.Println("insertId: " + string(insertId))
		writer.Write([]byte("ok"))
	}
}

// 获取所有 link
func getLinks(writer http.ResponseWriter, request *http.Request) {
	// 开始查询
	links, err := data.SelectLinks()
	if err != nil {
		writer.Write([]byte(err.Error()))
	} else {
		spew.Dump(links)
		writer.Write([]byte("ok"))
	}
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
	if err != nil {
		writer.Write([]byte(err.Error()))
	} else {
		spew.Dump(links)
		writer.Write([]byte("ok"))
	}
}
