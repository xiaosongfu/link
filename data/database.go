/*
 *
 * 测试服务器：192.168.160.3
 * 生产服务器：120.77.47.141
 *
 */
package data

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "root:link@(120.77.47.141:43306)/link?charset=utf8")
	if err != nil {
		panic(err)
	}

	if err := Db.Ping(); err != nil {
		panic(err)
	}

	log.Println("connect database success!")
}
