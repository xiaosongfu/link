package util

import "testing"

//func TestParseLinkTitle(t *testing.T) {
//	urls := []string{
//		//"http://www.flysnow.org/2018/02/09/go-regexp-extract-text.html",
//		"https://juejin.im/post/5addcfd76fb9a07aa349df55",
//	}
//
//	titles := []string{
//		//"Go语言中使用正则提取匹配的字符串 | 飞雪无情的博客",
//		//"一款基于Vue和Go的桌面端管理star项目应用",
//		"欧翁",
//	}
//
//	for index, url := range urls {
//		title, err := ParseLinkTitle(url)
//		t.Log(title)
//		if err != nil {
//			t.Error(err)
//		} else {
//			if title != titles[index] {
//				t.Error("title is not correct")
//			}
//		}
//	}
//}

func TestQueryLinkTitle(t *testing.T) {
	urls := []string{
		"http://www.flysnow.org/2018/02/09/go-regexp-extract-text.html",
		"https://juejin.im/post/5addcfd76fb9a07aa349df55",
	}

	titles := []string{
		"Go语言中使用正则提取匹配的字符串 | 飞雪无情的博客",
		"一款基于Vue和Go的桌面端管理star项目应用 - 掘金",
	}

	for index, url := range urls {
		title, err := QueryLinkTitle(url)
		t.Log(title)
		if err != nil {
			t.Error(err)
		} else {
			if title != titles[index] {
				t.Error("title is not correct")
			}
		}
	}
}

func TestIsCorrectUrl(t *testing.T) {
	urls := []string{
		"http://www.flysnow.org/2018/02/09/go-regexp-extract-text.html",
		"https://juejin.im/post/5addcfd76fb9a07aa349df55",
		"juejin.im/post/5addcfd76fb9a07aa349df55",
		"",
	}

	matchs := []bool{
		true,
		true,
		false,
		false,
	}

	for index, url := range urls {
		match := IsCorrectUrl(url)
		if match != matchs[index] {
			t.Error("is not correct url")
		}
	}
}

func TestStringIsEmpty(t *testing.T) {
	strs := []string{
		"",
		"asb",
		"是",
	}

	res := []bool{
		true,
		false,
		false,
	}

	for index, str := range strs {
		if res[index] != StringIsEmpty(str) {
			t.Error("wrong!")
		}
	}
}
