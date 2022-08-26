package routes

import (
	"app/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupAPI(api fiber.Router) {
	registerCountries(api)
}

func registerCountries(api fiber.Router) {
	countries := api.Group("/countries")

	countries.Get("/", handlers.GetAllContinents)
}
