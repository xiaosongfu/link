package main

import (
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

func QueryLinkTitle(url string) (title string, err error) {
	title = UnNamedTitle
	if url == "" {
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
