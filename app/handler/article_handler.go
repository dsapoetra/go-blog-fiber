package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go-blog-fiber/app/model/http"
	"go-blog-fiber/app/service"
	"go-blog-fiber/pkg/utils"
	"log"
	"strings"
)

func NewArticleHandler(app fiber.Router, userSrv service.IArticleService) {
	app.Get("/article/:id", GetArticle(userSrv))
	app.Post("/article", CreateArticle(userSrv))
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

// CreateArticle func gets author by given ID or 404 error.
// @Description Create article by given Body.
// @Summary get author by given ID
// @Tags Article
// @Accept json
// @Produce json
// @Param author body http.CreateArticleRequest true "Article"
// @Success 200 {object} repo.Article
// @Param Authorization header string true "Bearer"
// @Security ApiKeyAuth
// @Router /v1/article [post]
func CreateArticle(articleSrv service.IArticleService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Catch book ID from URL.
		body := &http.CreateArticleRequest{}

		if err := c.BodyParser(body); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":  true,
				"lalala": true,
				"msg":    err.Error(),
			})
		}

		header := c.GetReqHeaders()

		token := header["Authorization"]

		tokenArr := strings.Split(token, " ")

		t, err := utils.ValidateToken(tokenArr[1])

		log.Println(err)

		body.Author, _ = uuid.Parse(t.Id)

		err = articleSrv.CreateArticle(body)
		if err != nil {
			// Return, if book not found.
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"msg":   "error when inputting article",
			})
		}

		// Return status 200 OK.
		return c.JSON(fiber.Map{
			"error": false,
			"msg":   nil,
		})
	}
}
