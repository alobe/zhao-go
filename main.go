package main

import (
	"github.com/joho/godotenv"
	"zhao-go/lib/database"
	"zhao-go/lib/router"
)

func main() {
	godotenv.Load()
	// 路由注册
	router.InitRouter()
	//数据库初始化
	database.ConnectMysql()
}