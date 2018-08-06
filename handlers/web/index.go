package web

import (
	"html/template"
	"net/http"

	"go.xiaosongfu.com/link/handlers"
	"go.xiaosongfu.com/link/service"
)

// 注册路由
func RegisterIndexRoutes() {
	http.HandleFunc("/", handlers.Logger(Index))
}

// ----------------------------------------------------------------

func Index(writer http.ResponseWriter, request *http.Request) {
	// 开始查询
	links, err := service.GetLinks()
	if err != nil {
		panic(err)
	}

	temp, err := template.ParseFiles("web/templates/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(writer, links)
}
