package models

import (
	"fmt"
	"time"
)

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
