package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go-blog-fiber/app/model"
	"go-blog-fiber/app/service"
	"log"
	"time"
)

func NewAuthorHandler(app fiber.Router, authorSrv service.IAuthorService) {
	app.Get("/author/:id", GetAuthor(authorSrv))
	app.Post("/author", CreateAuthor(authorSrv))
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

		author, err := authorService.FindOneAuthor(id)
		if err != nil {
			// Return, if book not found.
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": true,
				"msg":   "author with the given ID is not found",
				"book":  nil,
			})
		}

		// Return status 200 OK.
		return c.JSON(fiber.Map{
			"error":  false,
			"msg":    nil,
			"author": author,
		})
	}
}

// CreateAuthor func gets author by given ID or 404 error.
// @Description Create author by given Body.
// @Summary get author by given ID
// @Tags Author
// @Accept json
// @Produce json
// @Param author body model.Author true "Author"
// @Success 200 {object} model.Author
// @Router /v1/author [post]
func CreateAuthor(authorService service.IAuthorService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Catch book ID from URL.
		body := &model.Author{}

		if err := c.BodyParser(body); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":  true,
				"lalala": true,
				"msg":    err.Error(),
			})
		}

		body.CreatedAt = time.Now()
		body.UpdatedAt = time.Now()

		log.Println(body.FullName)

		err := authorService.CreateAuthor(body)
		if err != nil {
			// Return, if book not found.
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"msg":   "error when inputting author",
			})
		}

		// Return status 200 OK.
		return c.JSON(fiber.Map{
			"error": false,
			"msg":   nil,
		})
	}
}
