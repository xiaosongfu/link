package handlers

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"github.com/xiaosongfu/link/data"
	"encoding/json"
)

func TestGetLinksByCategoryId(t *testing.T) {
	// 创建一个请求
	request := httptest.NewRequest(http.MethodGet, "categoryId=1", nil)
	//request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 创建一个 ResponseRecorder 来记录响应
	responseRecorder := httptest.NewRecorder()

	// 直接调用 RootHandler
	GetLinksByCategoryId(responseRecorder, request)

	// 检测返回的状态码
	if responseRecorder.Code != http.StatusOK {
		t.Errorf("response code is %v", responseRecorder.Code)
	}

	// 检测返回的数据
	var res = data.BasicResponse{}

	json.Unmarshal(responseRecorder.Body.Bytes(), &res)
	t.Log(res)
	if res.Code != 1 {
		t.Errorf("handler returned unexpected code: got %d want %d",
			res.Code, 1)
	}
}