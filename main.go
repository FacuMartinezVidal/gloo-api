package main

import (
	"fmt"
	_ "gloo-api/docs"
	"os"

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

	// Get port from environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Update Swagger host based on PORT
	host := fmt.Sprintf("localhost:%s", port)
	if os.Getenv("RAILWAY_STATIC_URL") != "" {
		host = os.Getenv("RAILWAY_STATIC_URL")
	}

	// Swagger configuration
	app.Static("/swagger", "./docs")
	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL: fmt.Sprintf("http://%s/swagger/swagger.json", host),
		DeepLinking: false,
		DocExpansion: "none",
	}))

	api := app.Group("/api")
	api.Get("/", helloWorld)

	app.Listen(":" + port)
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
