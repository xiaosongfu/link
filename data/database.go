package data

import (
	"database/sql"

	"log"

	_ "github.com/go-sql-driver/mysql"
	"ffll.fun/link/config"
)

var Db *sql.DB

func init() {
	log.Println("database connecting ...")

	database := config.Conf.Database[config.Conf.Env]
	var err error
	//Db, err = sql.Open("mysql", "root:link@(192.168.160.3:43306)/link?charset=utf8")
	Db, err = sql.Open(database.DriverName, database.UserName+":"+database.Password+"@("+database.Host+":"+database.Port+")/"+database.Database+"?charset=utf8")
	if err != nil {
		panic(err)
	}

	if err := Db.Ping(); err != nil {
		panic(err)
	}

	log.Println("database connect success!")
}
