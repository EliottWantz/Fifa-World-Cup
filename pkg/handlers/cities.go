package handlers

import (
	"app/pkg/db"
	"app/pkg/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllCities(c *fiber.Ctx) error {
	var cities []models.City
	db.Db.Find(&cities)
	return c.JSON(cities)
}

func GetCity(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).Send([]byte(err.Error()))
	}
	var city models.City
	if err := db.Db.First(&city, id).Error; err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.JSON(city)
}
