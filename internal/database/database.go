package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	seedDB(db)
	return db
}
