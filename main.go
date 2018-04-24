package main

import "net/http"

func main() {
	// 配置路由
	// --> category
	http.HandleFunc("/api/v1/addCategory", addCategory)
	http.HandleFunc("/api/v1/getCategorys", getCategorys)
	// --> link
	http.HandleFunc("/api/v1/addLink", addLink)
	http.HandleFunc("/api/v1/getLinks", getLinks)
	http.HandleFunc("/api/v1/getLinksByCategoryId", getLinksByCategoryId)

	// 启动服务
	err := http.ListenAndServe(":1205", nil)
	if err != nil {
		panic(err)
	}
}
