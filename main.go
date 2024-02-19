package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/kauri646/be-idealibs/config"

	"github.com/kauri646/be-idealibs/migration"
	"github.com/kauri646/be-idealibs/routes"
)

func main() {

	config.DatabaseInit()
	migration.RunMigration()

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	routes.RouteInit(app)

	app.Listen(":8080")

}
