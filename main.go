package main

import (
	"net/http"
	"github.com/xiaosongfu/link/route"
)

func main() {
	// 配置路由
	// --> category
	http.HandleFunc("/api/v1/addCategory", logger(route.AddCategory))
	http.HandleFunc("/api/v1/getCategorys", logger(route.GetCategorys))
	// --> link
	http.HandleFunc("/api/v1/addLink", logger(route.AddLink))
	http.HandleFunc("/api/v1/getLinks", logger(route.GetLinks))
	http.HandleFunc("/api/v1/getLinksByCategoryId", logger(route.GetLinksByCategoryId))

	// 启动服务
	err := http.ListenAndServe(":1205", nil)
	if err != nil {
		panic(err)
	}
}
