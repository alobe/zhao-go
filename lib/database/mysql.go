package database

import (
	"fmt"
	"log"
	"os"
	"zhao-go/lib/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectMysql() {
	dsn := os.Getenv("DATABASE_DSN")
	fmt.Println(dsn)
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
