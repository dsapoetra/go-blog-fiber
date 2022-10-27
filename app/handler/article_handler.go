package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go-blog-fiber/app/service"
)

func NewArticleHandler(app fiber.Router, userSrv service.IArticleService) {
	app.Get("/article/:id", GetArticle(userSrv))
}

// GetArticle func gets article by given ID or 404 error.
// @Description Get article by given ID.
// @Summary get article by given ID
// @Tags Article
// @Accept json
// @Produce json
// @Param id path string true "Article ID"
// @Success 200 {object} repo.Article
// @Router /v1/article/{id} [get]
func GetArticle(articleService service.IArticleService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Catch book ID from URL.
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}

		book, err := articleService.FindOneArticle(id)
		if err != nil {
			// Return, if book not found.
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": true,
				"msg":   "book with the given ID is not found",
				"book":  nil,
			})
		}

		// Return status 200 OK.
		return c.JSON(fiber.Map{
			"error": false,
			"msg":   nil,
			"book":  book,
		})
	}
}
