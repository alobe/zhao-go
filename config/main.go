package config

import (
	"time"

	"github.com/jinzhu/configor"
	"github.com/rs/zerolog/log"
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

	Jwt struct {
		Secret string
		Issuer string
		Expire time.Duration
	}
}{}

func InitConfig() {
	// db config load
	configor.Load(&Config, "./config/common.yaml")
	log.Info().Msgf("config ===> %v", Config)
}
