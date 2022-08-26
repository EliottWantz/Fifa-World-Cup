package models

import "time"

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
	ContinentID uint   `gorm:"not null" json:"continentId"`
}

type City struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	Name      string `gorm:"not null" json:"name"`
	Key       string `gorm:"not null" json:"key"`
	PlaceID   uint   `gorm:"not null" json:"placeId"`
	CountryID uint   `gorm:"not null" json:"countryId"`
}

type Event struct {
	ID       uint      `gorm:"primarykey" json:"id"`
	Key      string    `gorm:"not null" json:"key"`
	LeagueID uint      `gorm:"not null" json:"leagueId"`
	SeasonID uint      `gorm:"not null" json:"seasonId"`
	StartAt  time.Time `gorm:"not null" json:"startAt"`
}

type Group struct {
	ID      uint   `gorm:"primarykey" json:"id"`
	EventID uint   `gorm:"not null" json:"eventId"`
	Title   string `gorm:"not null" json:"title"`
	Pos     uint   `gorm:"not null" json:"pos"`
}

type Round struct {
	ID      uint      `gorm:"primarykey" json:"id"`
	EventID uint      `gorm:"not null" json:"eventId"`
	Title   string    `gorm:"not null" json:"title"`
	Pos     uint      `gorm:"not null" json:"pos"`
	StartAt time.Time `gorm:"not null" json:"startAt"`
	EndAt   time.Time `gorm:"not null" json:"endAt"`
}

type Game struct {
	ID       uint      `gorm:"primarykey" json:"id"`
	RoundID  uint      `gorm:"not null" json:"roundId"`
	Pos      uint      `gorm:"not null" json:"pos"`
	GroupID  uint      `gorm:"not null" json:"groupId"`
	Team1ID  uint      `gorm:"not null" json:"team1Id"`
	Team2ID  uint      `gorm:"not null" json:"team2Id"`
	PlayAt   time.Time `gorm:"not null" json:"playAt"`
	Score1   uint      `gorm:"not null" json:"score1"`
	Score2   uint      `gorm:"not null" json:"score2"`
	Score1ET uint      `json:"score1et"`
	Score2ET uint      `json:"score2et"`
	Score1P  uint      `json:"score1p"`
	Score2P  uint      ` json:"score2p"`
	Winner   uint      `json:"winner"`
	Winner90 uint      `json:"winner90"`
}

type Team struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	Key       string `gorm:"not null" json:"key"`
	Title     string `gorm:"not null" json:"title"`
	Code      string `json:"code"`
	Synonyms  string `json:"synonyms"`
	CountryID uint   `gorm:"not null" json:"countryId"`
}

type GroupTeam struct {
	ID      uint `gorm:"primarykey" json:"id"`
	GroupID uint `gorm:"not null" json:"groupId"`
	TeamID  uint `gorm:"not null" json:"teamId"`
}

type EventTeam struct {
	ID      uint `gorm:"primarykey" json:"id"`
	EventID uint `gorm:"not null" json:"eventId"`
	TeamID  uint `gorm:"not null" json:"teamId"`
}

type Person struct {
	ID   uint   `gorm:"primarykey" json:"id"`
	Key  string `gorm:"not null" json:"key"`
	Name string `gorm:"not null" json:"name"`
}

type Goal struct {
	ID       uint `gorm:"primarykey" json:"id"`
	PersonID uint `gorm:"not null" json:"personId"`
	GameID   uint `gorm:"not null" json:"gameId"`
	TeamID   uint `gorm:"not null" json:"teamId"`
	Minute   uint `gorm:"not null" json:"minute"`
	Offset   uint `gorm:"not null" json:"offset"`
	Score1   uint `gorm:"not null" json:"score1"`
	Score2   uint `gorm:"not null" json:"score2"`
}
