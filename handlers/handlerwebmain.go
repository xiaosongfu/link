package handlers

import (
	"ffll.fun/link/data"
	"html/template"
	"net/http"
)

func Index(writer http.ResponseWriter, request *http.Request) {
	// 开始查询
	links, err := data.SelectLinks()
	if err != nil {
		panic(err)
	}

	temp, err := template.ParseFiles("web/templates/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(writer, links)
}
