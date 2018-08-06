package dao

import "xiaosongfu.com/link/database"

// Category struct
//
// Category struct 的定义.
// 写点什么注释好呢
// swagger:model Category
type Category struct {
	// ID
	Id int `json:"id"`
	// 类别 ID
	CategoryId int `json:"categoryId"`
	// 类别名称
	CategoryName string `json:"categoryName"`
}

// 插入 category
func (category *Category) InsertCategory() (insertId int64, err error) {
	sqlStmt := "INSERT INTO category (category_id,category_name) VALUES (?,?)"
	stmt, err := database.Db.Prepare(sqlStmt)
	if err != nil {
		return
	}

	defer stmt.Close()

	result, err := stmt.Exec(category.CategoryId, category.CategoryName)
	if err != nil {
		return
	}
	insertId, err = result.LastInsertId()
	return
}

// 获取所有 category
func SelectCategorys() (categorys []Category, err error) {
	sqlStmt := "SELECT category_id,category_name FROM category"
	stmt, err := database.Db.Prepare(sqlStmt)
	if err != nil {
		return
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	defer rows.Close()

	for rows.Next() {
		category := Category{}
		if err = rows.Scan(&category.CategoryId, &category.CategoryName); err != nil {
			return
		}
		categorys = append(categorys, category)
	}
	return
}
