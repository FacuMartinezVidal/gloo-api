package main

import (
	_ "gloo-api/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	swagger "github.com/gofiber/swagger"
)

//	@title			Gloo API
//	@version		1.0
//	@description	API documentation for Gloo
//	@termsOfService	http://swagger.io/terms/
//	@contact.name	API Support
//	@contact.email	support@gloo.com
//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//	@host			localhost:8080
//	@BasePath		/api

func main() {
	app := fiber.New(fiber.Config{
		BodyLimit:             10 * 1024 * 1024, // 10MB
		ReadBufferSize:        16 * 1024 * 1024, // 16MB
		WriteBufferSize:       16 * 1024 * 1024, // 16MB
		DisableStartupMessage: false,
	})

	// CORS
	app.Use(cors.New())

	// Swagger configuration
	app.Static("/swagger", "./docs")
	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL: "http://localhost:8080/swagger/swagger.json",
		DeepLinking: false,
		DocExpansion: "none",
	}))

	api := app.Group("/api")
	api.Get("/", helloWorld)

	app.Listen(":8080")
}

// @Summary		Hello World endpoint
// @Description	Returns a simple hello world message
// @Tags			root
// @Accept			json
// @Produce		json
// @Success		200	{string}	string	"Hello, World!"
// @Router			/ [get]
func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
