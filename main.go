package main

import (
	"log"
	"net/http"

	"ffll.fun/link/config"
	"ffll.fun/link/handlers"
)

func main() {
	// 静态文件服务器
	fileServer := http.FileServer(http.Dir("web/public"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// 配置路由
	// --> 首页
	http.HandleFunc("/", handlers.Logger(handlers.Index))

	// --> category
	http.HandleFunc("/api/v1/addCategory", handlers.Logger(handlers.AddCategory))
	http.HandleFunc("/api/v1/getCategorys", handlers.Logger(handlers.GetCategorys))
	// --> link
	http.HandleFunc("/api/v1/addLink", handlers.Logger(handlers.AddLink))
	http.HandleFunc("/api/v1/getLinks", handlers.Logger(handlers.GetLinks))
	http.HandleFunc("/api/v1/getLinksByCategoryId", handlers.Logger(handlers.GetLinksByCategoryId))

	// 获取配置
	addr := config.Conf.Server[config.Conf.Env].Addr

	// 打印日志
	log.Println("server started,and listening on 0.0.0.0" + addr)

	// 启动服务
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		panic(err)
	}
}
