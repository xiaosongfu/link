package main

import "net/http"

func main() {
	// 配置路由
	http.HandleFunc("/api/v1/addLink", addLink)
	http.HandleFunc("/api/v1/getLinks", getLinks)
	http.HandleFunc("/api/v1/getLinksByCategoryId", getLinksByCategoryId)

	// 启动服务
	http.ListenAndServe(":1205", nil)
}
