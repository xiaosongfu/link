package dao

import "xiaosongfu.com/link/database"

// Link struct
//
// Link struct 的定义.
// 写点什么注释好呢
// swagger:model Link
type Link struct {
	// ID
	Id int `json:"id"`
	// 网页链接
	Url string `json:"url"`
	// 网页标题
	Title string `json:"title"`
	// 类别 ID
	CategoryId int `json:"categoryId"`
	// 标签
	Tag string `json:"tag"`
}

// category_id 默认值为"未分类",即是 1
const DefaultCategoryId int = 1

// 添加 link
func (link *Link) InsertLink() (insertId int64, err error) {
	sqlStmt := "INSERT INTO link (url,title,category_id,tag) VALUES (?,?,?,?)"
	stmt, err := database.Db.Prepare(sqlStmt)
	if err != nil {
		return
	}

	defer stmt.Close()

	result, err := stmt.Exec(link.Url, link.Title, link.CategoryId, link.Tag)
	if err != nil {
		return
	}

	insertId, err = result.LastInsertId()
	return
}

// 获取所有 link
func SelectLinks() (links []Link, err error) {
	sqlStmt := "SELECT id,url,title,category_id,tag FROM link ORDER BY id DESC"
	stmt, err := database.Db.Prepare(sqlStmt)
	if err != nil {
		return
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	defer rows.Close()

	for rows.Next() {
		link := Link{}
		if err = rows.Scan(&link.Id, &link.Url, &link.Title, &link.CategoryId, &link.Tag); err != nil {
			return
		}
		links = append(links, link)
	}
	return
}

// 根据 categoryId 获取 link
func SelectLinksByCategoryId(categoryId int) (links []Link, err error) {
	sqlStmt := "SELECT id,url,title,category_id,tag FROM link WHERE category_id=? ORDER BY id DESC"
	stmt, err := database.Db.Prepare(sqlStmt)
	if err != nil {
		return
	}

	defer stmt.Close()

	rows, err := stmt.Query(categoryId)

	defer rows.Close()

	for rows.Next() {
		link := Link{}
		if err = rows.Scan(&link.Id, &link.Url, &link.Title, &link.CategoryId, &link.Tag); err != nil {
			return
		}
		links = append(links, link)
	}
	return
}
