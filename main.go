package main

import (
	"zhao-go/config"
	"zhao-go/lib/database"
	"zhao-go/lib/logger"
	"zhao-go/lib/router"
)

func main() {
	// 初始化logger
	logger.InitialLog()
	// 初始化数据库配置
	config.InitConfig()
	//数据库初始化
	database.InitDB()
	// 路由注册
	router.InitRouter()
}
