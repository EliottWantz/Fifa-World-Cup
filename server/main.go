package main

import (
	"app/db"
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"app/routes"
)

var (
	port = flag.String("port", ":8080", "Port to listen on")
	prod = flag.Bool("prod", false, "Enable prefork in Production")
)

func main() {
	// Parse command-line flags
	flag.Parse()

	// Connected with database
	db.Connect()

	// Create fiber app
	app := fiber.New(fiber.Config{
		Prefork: *prod, // go run app.go -prod
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())

	// Create a /api/v1 endpoint
	v1 := app.Group("/api/v1")

	// Setup routes
	routes.SetupAPI(v1)

	// Setup static files
	app.Static("/", "../client/dist")

	// Listen on port 8080
	log.Fatal(app.Listen(*port)) // go run app.go -port=:8080
}
