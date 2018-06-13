package web

import (
	"html/template"
	"net/http"

	"ffll.fun/link/service"
)

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
