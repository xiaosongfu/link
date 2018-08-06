// Copyright 2018 xiaosong fu. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//     link 项目 swagger 配置
//
//     该文档描述了 link 项目所有的 api 入参，出参配置
//
//     Schemes: http
//     Host: localhost:1205
//     BasePath: /
//     Version: 1.0.0
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: xiaosong fu<xiaosong.fu@outlook.com> https://notes.fuxiaosong.cn
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
	"log"
	"net/http"

	"go.xiaosongfu.com/link/config"
	"go.xiaosongfu.com/link/handlers/api"
	"go.xiaosongfu.com/link/handlers/web"
	_ "go.xiaosongfu.com/link/swagger"
)

func main() {
	// 静态文件服务器
	fileServer := http.FileServer(http.Dir("web/public"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// 配置路由
	web.RegisterIndexRoutes()    // --> 首页
	api.RegisterCategoryRoutes() // --> category
	api.RegisterLinkRoutes()     // --> link
	//// --> 首页
	//http.HandleFunc("/", handlers.Logger(web.Index))
	//
	//// --> category
	//http.HandleFunc("/api/v1/addCategory", handlers.Logger(api.AddCategory))
	//http.HandleFunc("/api/v1/getCategorys", handlers.Logger(api.GetCategorys))
	//// --> link
	//http.HandleFunc("/api/v1/addLink", handlers.Logger(api.AddLink))
	//http.HandleFunc("/api/v1/getLinks", handlers.Logger(api.GetLinks))
	//http.HandleFunc("/api/v1/getLinksByCategoryId", handlers.Logger(api.GetLinksByCategoryId))

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
