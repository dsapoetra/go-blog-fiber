package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-blog-fiber/app/handler"
)

func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/v1")

	// Routes for GET method:
	route.Get("/article/:id", handler.GetArticle) // get one book by ID
}
