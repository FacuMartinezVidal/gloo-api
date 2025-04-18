package main

import (
	_ "gloo-api/docs"
	"gloo-api/routes"
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
//	@host			gloo-api-production.up.railway.app
//	@BasePath		/api
//	@schemes		https

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

	// Swagger configuration
	app.Static("/swagger", "./docs")
	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL: "https://gloo-api-production.up.railway.app/swagger/swagger.json",
		DeepLinking: false,
		DocExpansion: "none",
	}))

	app.Get("/recipes/:searchParam/:category", routes.GetAllRecipe)


	app.Listen(":" + port)
}


