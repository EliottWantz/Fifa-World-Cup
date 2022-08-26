package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("worldcup.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	for _, c := range getAllContinents(db) {
		c.Print()
	}
}
