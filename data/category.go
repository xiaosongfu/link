package data

// Category struct
type Category struct {
	Id           int    `json:"id"`
	CategoryId   int    `json:"categoryId"`
	CategoryName string `json:"categoryName"`
}

// 插入 category
func (category *Category) InsertCategory() (insertId int64, err error) {
	sqlStmt := "INSERT INTO category (category_id,category_name) VALUES (?,?)"
	stmt, err := Db.Prepare(sqlStmt)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	result, err := stmt.Exec(category.CategoryId, category.CategoryName)
	if err != nil {
		panic(err)
	}
	insertId, err = result.LastInsertId()
	return
}

// 获取所有 category
func SelectCategorys() (categorys []Category, err error) {
	sqlStmt := "SELECT category_id,category_name FROM category"
	stmt, err := Db.Prepare(sqlStmt)
	if err != nil {
		panic(err)
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
