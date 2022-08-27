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
	registerRounds(api)
	registerTeams(api)
	registerGoals(api)
	registerScorers(api)
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

func registerRounds(api fiber.Router) {
	rounds := api.Group("/rounds")

	rounds.Get("/", handlers.GetAllRounds)
	rounds.Get("/:id", handlers.GetRound)
}

func registerTeams(api fiber.Router) {
	teams := api.Group("/teams")

	teams.Get("/", handlers.GetAllTeams)
	teams.Get("/:id", handlers.GetTeam)
}

func registerGoals(api fiber.Router) {
	goals := api.Group("/goals")

	goals.Get("/", handlers.GetAllGoals)
	goals.Get("/:id", handlers.GetGoal)
}

func registerScorers(api fiber.Router) {
	scorers := api.Group("/scorers")

	scorers.Get("/", handlers.GetAllScorers)
	scorers.Get("/:id", handlers.GetScorer)
}
