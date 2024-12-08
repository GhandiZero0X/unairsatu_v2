package main

import (
	"project-crud_baru/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	routes.RouteApp(app)
	app.Listen(":3000")
}
