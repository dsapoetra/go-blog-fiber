package main

import (
	"github.com/gofiber/fiber/v2"
	"go-blog-fiber/pkg/configs"
	"go-blog-fiber/pkg/utils"

	//_ "fiber/docs"                        // load API Docs files (Swagger)
	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

func main() {
	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	//// Middlewares.
	//middleware.FiberMiddleware(app) // Register Fiber's middleware for app.
	//
	//// Routes.
	//routes.SwaggerRoute(app)  // Register a route for API Docs (Swagger).
	//routes.PublicRoutes(app)  // Register a public routes for app.
	//routes.PrivateRoutes(app) // Register a private routes for app.
	//routes.NotFoundRoute(app) // Register route for 404 Error.

	// Start server (with graceful shutdown).
	utils.StartServerWithGracefulShutdown(app)
}
