package setup

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(env *Env) *gorm.DB {
	log.Println("Connect to database dns: ", env.DatabaseDsn)
	db, err := gorm.Open(postgres.Open(env.DatabaseDsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection error: ", err.Error())
	}

	return db
}
