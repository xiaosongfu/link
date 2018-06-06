package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Env      string
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
	log.Println("start read config info ...")

	if _, err := toml.DecodeFile(configFile, &Conf); err != nil {
		panic(err)
	}

	log.Println("read config info success.")
}
