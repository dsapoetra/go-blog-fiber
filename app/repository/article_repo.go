package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"go-blog-fiber/app/model"
	"log"
)

type ArticleRepository struct {
	db *sqlx.DB
}

type IArticleRepository interface {
	FindOneArticle(id uuid.UUID) (*model.Article, error)
}

func NewArticleRepository(db *sqlx.DB) IArticleRepository {
	return &ArticleRepository{
		db: db,
	}
}

func (a *ArticleRepository) FindOneArticle(id uuid.UUID) (*model.Article, error) {
	article := model.Article{}
	log.Println("HEREEEEEEEEEE 2")

	query := `SELECT * FROM articles where id = $1`

	err := a.db.Get(&article, query, id)

	if err != nil {
		log.Println("HEREEEEEEEEEE")
		return nil, err
	}

	return &article, nil
}
