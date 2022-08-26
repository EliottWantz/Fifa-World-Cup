package models

type Continent struct {
	ID      uint   `gorm:"primarykey" json:"id"`
	Name    string `gorm:"not null" json:"name"`
	Slug    string `gorm:"not null" json:"slug"`
	Key     string `gorm:"not null" json:"key"`
	PlaceID uint   `gorm:"not null" json:"placeId"`
}

type Country struct {
	ID          uint   `gorm:"primarykey" json:"id"`
	Name        string `gorm:"not null" json:"name"`
	Slug        string `gorm:"not null" json:"slug"`
	Key         string `gorm:"not null" json:"key"`
	PlaceID     uint   `gorm:"not null" json:"placeId"`
	Code        string `gorm:"not null" json:"code"`
	Pop         uint   `gorm:"not null" json:"pop"`
	Area        uint   `gorm:"not null" json:"area"`
	ContinentId uint   `gorm:"not null" json:"continentId"`
}

type City struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	Name      string `gorm:"not null" json:"name"`
	Key       string `gorm:"not null" json:"key"`
	PlaceID   uint   `gorm:"not null" json:"placeId"`
	CountryId uint   `gorm:"not null" json:"countryId"`
}
