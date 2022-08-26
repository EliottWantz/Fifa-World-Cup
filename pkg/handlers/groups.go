package handlers

import (
	"app/pkg/db"
	"app/pkg/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllGroups(c *fiber.Ctx) error {
	var groups []models.Group
	db.Db.Where(&models.Group{EventID: 21}).Find(&groups)
	return c.JSON(groups)
}

func GetGroup(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).Send([]byte("id must be a number"))
	}
	var group models.Group
	if err := db.Db.Where(&models.Group{EventID: 21}).First(&group, id).Error; err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.JSON(group)
}
