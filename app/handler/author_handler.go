package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go-blog-fiber/app/service"
)

func NewAuthorHandler(app fiber.Router, authorSrv service.IAuthorService) {
	app.Get("/author/:id", GetAuthor(authorSrv))
}

// GetAuthor func gets author by given ID or 404 error.
// @Description Get author by given ID.
// @Summary get author by given ID
// @Tags Author
// @Accept json
// @Produce json
// @Param id path string true "Author ID"
// @Success 200 {object} model.Author
// @Router /v1/author/{id} [get]
func GetAuthor(authorService service.IAuthorService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Catch book ID from URL.
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}

		book, err := authorService.FindOneAuthor(id)
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
