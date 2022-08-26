package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	Db *gorm.DB
)

func Connect() {
	var err error
	Db, err = gorm.Open(sqlite.Open("./pkg/db/worldcup.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	log.Println("Connected to db")
	Db.Logger = logger.Default.LogMode(logger.Info)
}
