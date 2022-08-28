package config

import (
	"fmt"

	"github.com/jinzhu/configor"
)

var Config = struct {
	Mysql struct {
		Dsn string
	}

	Redis struct {
		Address  string
		Password string
		Db       int
	}
}{}

func InitConfig() {
	// db config load
	configor.Load(&Config, "./config/db.yaml")
	fmt.Printf("config %#v", Config)
}
