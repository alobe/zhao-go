package database

import (
	"log"
	"zhao-go/config"
	"zhao-go/lib/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectMysql() {
	dsn := config.Config.Mysql.Dsn
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	log.Print(db)
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&model.Image{})
	if err != nil {
		log.Fatal(err)
	}
	DB = db
}

func GetTableDb(model interface{}, table string) *gorm.DB {
	return DB.Model(model).Table(table)
}
