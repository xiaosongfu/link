package service

import (
	"ffll.fun/link/dao"
)

// 添加 link
func AddLink(link *dao.Link) (insertId int64, err error) {
	return link.InsertLink()
}

// 获取所有 link
func GetLinks() (links []dao.Link, err error) {
	return dao.SelectLinks()
}

// 根据 categoryId 获取 link
func GetLinksByCategoryId(categoryId int) (links []dao.Link, err error) {
	return dao.SelectLinksByCategoryId(categoryId)
}
