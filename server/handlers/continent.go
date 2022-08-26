package handlers

import (
	"app/db"
	"app/models"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllContinents(c *fiber.Ctx) error {
	var continents []models.Continent
	db.DB.Find(&continents)
	return c.JSON(continents)
}

func GetContinent(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fmt.Errorf("%v : %v", fiber.ErrBadRequest, err)
	}
	var continent models.Continent
	db.DB.First(&continent, id)
	return c.JSON(continent)
}
