package main

import (
	"notcoder_exe__blog/database"
	"notcoder_exe__blog/handlers"

	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var (
	port = flag.String("port", ":3000", "Port to listen on")
	prod = flag.Bool("prod", false, "Enable prefork in Production")
)

func main() {
	// Parse command-line flags
	flag.Parse()

	// Connected with database
	database.Connect()

	// Create fiber app
	app := fiber.New(fiber.Config{
		Prefork: *prod, // go run app.go -prod
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())

	// Create a /api/v1 endpoint
	v1 := app.Group("/api/v1")

	// Bind handlers
	v1.Route("/usertype", func(router fiber.Router) {
		router.Get("/", handlers.UserTypeGet)
		router.Get("/:id", handlers.UserTypeGetByID)
		router.Post("/", handlers.UserTypeCreate)
		router.Delete("/:id", handlers.UserTypeDelete)
		router.Put("/:id", handlers.UserTypeUpdate)
	})

	v1.Route("/field", func(router fiber.Router) {
		router.Post("/", handlers.FieldCreate)
		router.Get("/", handlers.FieldGet)
	})

	// Setup static files
	app.Static("/", "./static/public")

	// Handle not founds
	app.Use(handlers.NotFound)

	// Listen on port 3000
	log.Fatal(app.Listen(*port)) // go run app.go -port=:3000
}
