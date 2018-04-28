package util

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
)

const UnNamedTitle = "未命名"

//func ParseLinkTitle(url string) (title string, err error) {
//	title = UnNamedTitle
//	if url == "" {
//		err = errors.New("url 为空")
//		return
//	}
//
//	resp, err := http.Get(url)
//	if err != nil {
//		return
//	}
//
//	if resp.StatusCode == http.StatusOK {
//		var bytes []byte
//		bytes, err = ioutil.ReadAll(resp.Body)
//		if err != nil {
//			return
//		} else {
//			compile := regexp.MustCompile(`^<title>([.]+)</title>$`)
//			subMatch := compile.FindStringSubmatch(string(bytes))
//			if len(subMatch) >= 2 {
//				title = subMatch[1]
//			}
//		}
//	}
//	return
//}

// 根据网页的 url 解析出网页的 title
func QueryLinkTitle(url string) (title string, err error) {
	title = UnNamedTitle
	if StringIsEmpty(url) {
		err = errors.New("url 为空")
		return
	}
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return
	}
	title = doc.Find("title").Text()
	return
}

// 判断 url 是否合法
func IsCorrectUrl(url string) (match bool) {
	matchHttp, _ := regexp.MatchString("^http://", url)
	matchHttps, _ := regexp.MatchString("^https://", url)
	return matchHttp || matchHttps
}

// 判断字符串是否为空
func StringIsEmpty(str string) (empty bool) {
	if str == "" {
		empty = true
	} else {
		empty = false
	}
	return
}

// 向 http.ResponseWriter 写入 json 响应
func WriteJsonResponse(writer http.ResponseWriter, data interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	jsonResp, _ := json.Marshal(data)
	writer.Write(jsonResp)
}
