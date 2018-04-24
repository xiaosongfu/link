package data

// Link struct
type Link struct {
	Id         int    `json:"id"`
	Url        string `json:"url"`
	CategoryId int    `json:"categoryId"`
	Tag        string `json:"tag"`
}

// category_id 默认值为"未分类",即是 1
const DefaultCategoryId int = 1

// 添加 link
func (link *Link) InsertLink() (insertId int64, err error) {
	sqlStmt := "INSERT INTO link (url,category_id,tag) VALUES (?,?,?)"
	stmt, err := Db.Prepare(sqlStmt)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	result, err := stmt.Exec(link.Url, link.CategoryId, link.Tag)
	if err != nil {
		panic(err)
	}

	insertId, err = result.LastInsertId()
	return
}

// 获取所有 link
func SelectLinks() (links []Link, err error) {
	sqlStmt := "SELECT id,url,category_id,tag FROM link"
	stmt, err := Db.Prepare(sqlStmt)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	defer rows.Close()

	for rows.Next() {
		link := Link{}
		if err = rows.Scan(&link.Id, &link.Url, &link.CategoryId, &link.Tag); err != nil {
			return
		}
		links = append(links, link)
	}
	return
}

// 根据 categoryId 获取 link
func SelectLinksByCategoryId(categoryId int) (links []Link, err error) {
	sqlStmt := "SELECT id,url,category_id,tag FROM link WHERE category_id=?"
	stmt, err := Db.Prepare(sqlStmt)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query(categoryId)

	defer rows.Close()

	for rows.Next() {
		link := Link{}
		if err = rows.Scan(&link.Id, &link.Url, &link.CategoryId, &link.Tag); err != nil {
			return
		}
		links = append(links, link)
	}
	return
}