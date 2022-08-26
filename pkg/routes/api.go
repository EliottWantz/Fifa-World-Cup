package routes

import (
	"app/pkg/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupAPI(api fiber.Router) {
	registerContinents(api)
	registerCountries(api)
	registerCities(api)
	registerGroups(api)
}

func registerContinents(api fiber.Router) {
	continents := api.Group("/continents")

	continents.Get("/", handlers.GetAllContinents)
	continents.Get("/:id", handlers.GetContinent)
}

func registerCountries(api fiber.Router) {
	countries := api.Group("/countries")

	countries.Get("/", handlers.GetAllCountries)
	countries.Get("/:id", handlers.GetCountry)
}

func registerCities(api fiber.Router) {
	cities := api.Group("/cities")

	cities.Get("/", handlers.GetAllCities)
	cities.Get("/:id", handlers.GetCity)
}

func registerGroups(api fiber.Router) {
	groups := api.Group("/groups")

	groups.Get("/", handlers.GetAllGroups)
	groups.Get("/:id", handlers.GetGroup)
}
