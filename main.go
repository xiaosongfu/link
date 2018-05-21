package main

import (
	"net/http"

	"github.com/golang/glog"
	"github.com/xiaosongfu/link/config"
	"github.com/xiaosongfu/link/handlers"
)

func main() {
	glog.Infoln("server starting...")

	// 配置路由
	// --> category
	http.HandleFunc("/api/v1/addCategory", handlers.Logger(handlers.AddCategory))
	http.HandleFunc("/api/v1/getCategorys", handlers.Logger(handlers.GetCategorys))
	// --> link
	http.HandleFunc("/api/v1/addLink", handlers.Logger(handlers.AddLink))
	http.HandleFunc("/api/v1/getLinks", handlers.Logger(handlers.GetLinks))
	http.HandleFunc("/api/v1/getLinksByCategoryId", handlers.Logger(handlers.GetLinksByCategoryId))

	// 启动服务
	addr := config.Conf.Server[config.Conf.Env].Addr
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		panic(err)
	}

	glog.Infoln("server started,and listening on " + addr)
}
