package service

import (
	"ffll.fun/link/dao"
)

// 插入 category
func AddCategory(category *dao.Category) (insertId int64, err error) {
	return category.InsertCategory()
}

// 获取所有 category
func GetCategorys() (categorys []dao.Category, err error) {
	return dao.SelectCategorys()
}
