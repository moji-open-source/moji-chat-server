package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var __db__ *gorm.DB

func InitDatabase() error {
	log.Println("init database...")

	dsn := "host=localhost user=postgres password=123456 dbname=moji_chat port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	__db__ = db

	return err
}

func GetDatabase() *gorm.DB {
	return __db__
}
