package handlers

import (
	"app/pkg/db"
	"app/pkg/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllContinents(c *fiber.Ctx) error {
	var continents []models.Continent
	db.Db.Find(&continents)
	return c.JSON(continents)
}

func GetContinent(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).Send([]byte(err.Error()))
	}
	var continent models.Continent
	if err := db.Db.First(&continent, id).Error; err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.JSON(continent)
}
