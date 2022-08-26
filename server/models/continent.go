package models

import (
	"fmt"
	"time"
)

type Continent struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	Slug      string    `gorm:"not null" json:"slug"`
	Key       string    `gorm:"not null" json:"key"`
	PlaceID   int       `gorm:"not null" json:"placeId"`
	AltNames  string    `gorm:"not null" json:"altNames"`
	CreatedAt time.Time `gorm:"not null" json:"createdAt"`
	UpdatedAt time.Time `gorm:"not null" json:"updatedAt"`
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
