/*
 *
 * 测试服务器：192.168.160.3
 * 生产服务器：120.77.47.141
 *
 */
package data

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"github.com/xiaosongfu/link/config"
)

var Db *sql.DB

func init() {
	glog.Infoln("database connecting ...")

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

	glog.Infoln("database connect success!")
}
