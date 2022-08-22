package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"zhao-go/lib/model"
)

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
}

var DB *gorm.DB
