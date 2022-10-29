package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"go-blog-fiber/app/model/repo"
)

type ArticleRepository struct {
	db *sqlx.DB
}

type IArticleRepository interface {
	FindOneArticle(id uuid.UUID) (*repo.Article, error)
}

func NewArticleRepository(db *sqlx.DB) IArticleRepository {
	return &ArticleRepository{
		db: db,
	}
}

func (a *ArticleRepository) FindOneArticle(id uuid.UUID) (*repo.Article, error) {
	article := repo.Article{}

	query := `SELECT * FROM articles where id = $1`

	err := a.db.Get(&article, query, id)

	if err != nil {
		return nil, err
	}

	return &article, nil
}
