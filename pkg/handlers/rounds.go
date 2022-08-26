package handlers

import (
	"app/pkg/db"
	"app/pkg/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllRounds(c *fiber.Ctx) error {
	var rounds []models.Round
	db.Db.Where(&models.Round{EventID: 21}).Find(&rounds)
	return c.JSON(rounds)
}

func GetRound(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).Send([]byte("id must be a number"))
	}
	var round models.Round
	if err := db.Db.Where(&models.Round{EventID: 21}).First(&round, id).Error; err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.JSON(round)
}
