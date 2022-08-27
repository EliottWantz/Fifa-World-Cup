package handlers

import (
	"app/pkg/db"
	"app/pkg/models"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Query struct {
	GameID      string `query:"game_id"`
	RoundID     string `query:"round_id"`
	GroupID     string `query:"group_id"`
	TeamID      string `query:"team_id"`
	CountryID   string `query:"country_id"`
	ContinentID string `query:"continent_id"`
}

func GetAllScorers(c *fiber.Ctx) error {
	var persons []models.Person
	query := new(Query)
	err := c.QueryParser(query)
	if err != nil {
		return err
	}
	fmt.Println("THE query: ", query)
	db_exec := db.Db.Select("persons.*").Table("persons").Joins("JOIN goals ON goals.person_id = persons.id JOIN games ON goals.game_id = games.id JOIN rounds ON games.round_id = rounds.id").Where("rounds.event_id = 21")
	if query.GameID != "" {
		db_exec.Where("games.id = ?", query.GameID)
	}
	if query.RoundID != "" {
		db_exec.Where("rounds.id = ?", query.RoundID)
	}
	if query.GroupID != "" {
		db_exec.Joins("JOIN groups ON groups.id = games.group_id").Where("groups.id = ?", query.GroupID)
	}
	if query.TeamID != "" {
		db_exec.Joins("JOIN teams ON goals.team_id = teams.id").Where("teams.id = ?", query.TeamID)
	}
	if query.CountryID != "" {
		db_exec.Joins("JOIN teams ON goals.team_id = teams.id JOIN countries ON teams.country_id = countries.id").Where("countries.id = ?", query.CountryID)
	}
	if query.ContinentID != "" {
		db_exec.Joins("JOIN teams ON goals.team_id = teams.id JOIN countries ON teams.country_id = countries.id JOIN continents ON continents.id = countries.continent_id").Where("continents.id = ?", query.ContinentID)
	}
	db_exec.Find(&persons)
	return c.JSON(persons)
}

func GetScorer(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).Send([]byte("id must be a number"))
	}
	var person models.Person
	if err := db.Db.Select("persons.*").Table("persons").Joins("JOIN goals ON goals.person_id = persons.id JOIN games ON goals.game_id = games.id JOIN rounds ON games.round_id = rounds.id").Where("rounds.event_id = 21 AND persons.id = ?", id).First(&person).Error; err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.JSON(person)
}
