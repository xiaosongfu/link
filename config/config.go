package config

import (
	"github.com/BurntSushi/toml"
	"github.com/golang/glog"
)

type Config struct {
	Env string
	Server   map[string]Server
	Database map[string]Database
}

type Server struct {
	Addr string
}

type Database struct {
	DriverName string
	Host       string
	Port       string
	Database   string
	UserName   string
	Password   string
}

const configFile = "config/config.toml"

var Conf Config

func init() {
	glog.Infoln("start read config info ...")

	if _, err := toml.DecodeFile(configFile, &Conf); err != nil {
		panic(err)
	}

	glog.Infoln("read config info success.")
}
