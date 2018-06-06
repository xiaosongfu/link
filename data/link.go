package data

// Link struct
type Link struct {
	Id         int    `json:"id"`
	Url        string `json:"url"`
	Title      string `json:"title"`
	CategoryId int    `json:"categoryId"`
	Tag        string `json:"tag"`
}

// category_id 默认值为"未分类",即是 1
const DefaultCategoryId int = 1

// 添加 link
func (link *Link) InsertLink() (insertId int64, err error) {
	sqlStmt := "INSERT INTO link (url,title,category_id,tag) VALUES (?,?,?,?)"
	stmt, err := Db.Prepare(sqlStmt)
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
	sqlStmt := "SELECT id,url,title,category_id,tag FROM link ORDER BY id ASC"
	stmt, err := Db.Prepare(sqlStmt)
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
	sqlStmt := "SELECT id,url,title,category_id,tag FROM link WHERE category_id=? ORDER BY id ASC"
	stmt, err := Db.Prepare(sqlStmt)
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
