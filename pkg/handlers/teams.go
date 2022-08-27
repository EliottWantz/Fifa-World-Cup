package handlers

import (
	"app/pkg/db"
	"app/pkg/models"

	"github.com/gofiber/fiber/v2"
)

func GetAllTeams(c *fiber.Ctx) error {
	var teams []models.Team
	db.Db.Select("teams.*").Joins("JOIN events_teams AS et ON et.team_id=teams.id").Where("et.event_id = 21").Find(&teams)
	return c.JSON(teams)
}

func GetTeam(c *fiber.Ctx) error {
	id := c.Params("id")
	var team models.Team
	if err := db.Db.Select("teams.*").Joins("JOIN events_teams AS et ON et.team_id=teams.id").Where("et.event_id = 21 AND teams.id = ?", id).First(&team).Error; err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.JSON(team)
}
