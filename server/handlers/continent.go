package handlers

import (
	"app/db"
	"app/models"

	"github.com/gofiber/fiber/v2"
)

func GetAllContinents(c *fiber.Ctx) error {
	var continents []models.Continent
	db.DB.Find(&continents)
	return c.JSON(continents)
}
