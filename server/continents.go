package main

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

// CREATE TABLE IF NOT EXISTS "continents" ("id" integer PRIMARY KEY AUTOINCREMENT NOT NULL, "name" varchar NOT NULL, "slug" varchar NOT NULL, "key" varchar NOT NULL, "place_id" integer NOT NULL, "alt_names" varchar, "created_at" datetime NOT NULL, "updated_at" datetime NOT NULL);
// CREATE INDEX "index_continents_on_place_id" ON "continents" ("place_id");
// CREATE UNIQUE INDEX "index_continents_on_key" ON "continents" ("key");

type Continent struct {
	ID        uint      `gorm:"primarykey"`
	Name      string    `gorm:"not null"`
	Slug      string    `gorm:"not null"`
	Key       string    `gorm:"not null"`
	PlaceID   int       `gorm:"not null"`
	AltNames  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

func (c *Continent) Print() {
	fmt.Println("ID:", c.ID)
	fmt.Println("Name:", c.Name)
	fmt.Println("Slug:", c.Slug)
	fmt.Println("Key:", c.Key)
	fmt.Println("PlaceID:", c.PlaceID)
	fmt.Println("AltNames:", c.AltNames)
	fmt.Println("CreatedAt:", c.CreatedAt)
	fmt.Println("UpdatedAt:", c.UpdatedAt)
}

func getAllContinents(db *gorm.DB) []Continent {
	var continents []Continent
	db.Find(&continents)
	return continents
}
