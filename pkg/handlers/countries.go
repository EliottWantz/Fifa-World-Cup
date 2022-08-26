package handlers

import (
	"app/pkg/db"
	"app/pkg/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllCountries(c *fiber.Ctx) error {
	var countries []models.Country
	db.Db.Find(&countries)
	return c.JSON(countries)
}

func GetCountry(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).Send([]byte(err.Error()))
	}
	var country models.Country
	if err := db.Db.First(&country, id).Error; err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.JSON(country)
}
