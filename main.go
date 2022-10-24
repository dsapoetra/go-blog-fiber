package main

import (
	"github.com/gofiber/fiber/v2"
	"go-blog-fiber/app/handler"
	"go-blog-fiber/app/repository"
	"go-blog-fiber/app/routes"
	"go-blog-fiber/app/service"
	"go-blog-fiber/pkg/configs"
	"go-blog-fiber/pkg/utils"
	"go-blog-fiber/platform/database"

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
	_ "go-blog-fiber/docs"                // load API Docs files (Swagger)
)

func main() {
	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	db, _ := database.PostgreSQLConnection()
	articleRepo := repository.NewArticleRepository(db)
	articleSvc := service.NewArticleService(articleRepo)

	authorRepo := repository.NewAuthorRepository(db)
	authorSvc := service.NewAuthorService(authorRepo)
	//// Middlewares.
	//middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	//// Routes.
	routes.SwaggerRoute(app) // Register a route for API Docs (Swagger).
	//routes.PrivateRoutes(app) // Register a private routes for app.
	//routes.NotFoundRoute(app) // Register route for 404 Error.

	api := app.Group("/v1")
	app.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.JSON("ok")
	})

	// Prepare our endpoints for the API.
	handler.NewArticleHandler(api.Group("/"), articleSvc)
	handler.NewAuthorHandler(api.Group("/"), authorSvc)

	// Start server (with graceful shutdown).
	utils.StartServerWithGracefulShutdown(app)
}
