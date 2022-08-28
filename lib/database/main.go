package database

func InitDB() {
	ConnectMysql()
	ConnectRedis()
}
