package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func NewDB() *gorm.DB {

	dsn := "host=db user=pguser password=pguser dbname=pguser port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect to db: %s", err.Error())
	}

	return db
}
