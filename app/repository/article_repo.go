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
	CreateArticle(article *repo.Article) error
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

func (a *ArticleRepository) CreateArticle(article *repo.Article) error {
	query := `INSERT INTO articles(created_at, updated_at, title, content, author, likes, is_deleted) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := a.db.Exec(query, article.CreatedAt, article.UpdatedAt, article.Title, article.Content, article.Author, 0, false)

	return err
}
