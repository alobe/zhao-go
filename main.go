package main

import (
	"zhao-go/config"
	"zhao-go/lib/database"
	"zhao-go/lib/logger"
	"zhao-go/lib/router"
	// "zhao-go/lib/util"
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
	// var images [] string
	// // path := "./config/image.json"
	// // Json := util.Json[struct {images []string}]{
	// // 	path: path,
	// // 	value: struct {images []string}{
	// // 		images: images,
	// // 	},
	// // }
}
