package handlers

import (
	"app/pkg/db"
	"app/pkg/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllGoals(c *fiber.Ctx) error {
	var goals []models.Goal
	game_id := c.Query("game_id")
	round_id := c.Query("round_id")
	db_exec := db.Db.Select("goals.*").Joins("JOIN games AS g ON goals.game_id = g.id").Joins("JOIN rounds AS r ON g.round_id = r.id").Where("r.event_id = 21")
	if game_id != "" {
		db_exec.Where("g.id = ?", game_id)
	}
	if round_id != "" {
		db_exec.Where("r.id = ?", round_id)
	}
	db_exec.Find(&goals)
	return c.JSON(goals)
}

func GetGoal(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).Send([]byte("id must be a number"))
	}
	var goal models.Goal
	if err := db.Db.Select("goals.*").Joins("JOIN games AS g ON goals.game_id = g.id").Joins("JOIN rounds AS r ON g.round_id = r.id").Where("r.event_id = 21 AND goals.id = ?", id).First(&goal).Error; err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.JSON(goal)
}
