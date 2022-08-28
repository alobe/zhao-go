package database

import (
	"zhao-go/config"

	"github.com/go-redis/redis/v8"
)

var Redis *redis.Client

func ConnectRedis() {
	cfg := config.Config.Redis

	Redis = redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password,
		DB:       cfg.Db,
	})
}
