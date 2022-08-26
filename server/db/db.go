package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Connect() {
	db, err := gorm.Open(sqlite.Open("/db/worldcup.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
}
