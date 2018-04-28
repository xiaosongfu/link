package main

import (
	"net/http"
	"github.com/xiaosongfu/link/route"
)

func main() {
	// 配置路由
	// --> category
	http.HandleFunc("/api/v1/addCategory", route.AddCategory)
	http.HandleFunc("/api/v1/getCategorys", route.GetCategorys)
	// --> link
	http.HandleFunc("/api/v1/addLink", route.AddLink)
	http.HandleFunc("/api/v1/getLinks", route.GetLinks)
	http.HandleFunc("/api/v1/getLinksByCategoryId", route.GetLinksByCategoryId)

	// 启动服务
	err := http.ListenAndServe(":1205", nil)
	if err != nil {
		panic(err)
	}
}
