package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"go-blog-fiber/app/model"
)

type AuthorRepository struct {
	db *sqlx.DB
}

type IAuthorRepository interface {
	FindOneAuthor(id uuid.UUID) (*model.Author, error)
}

func NewAuthorRepository(db *sqlx.DB) IAuthorRepository {
	return &AuthorRepository{
		db: db,
	}
}

func (a *AuthorRepository) FindOneAuthor(id uuid.UUID) (*model.Author, error) {
	author := model.Author{}

	query := `SELECT * FROM authors where id = $1`

	err := a.db.Get(&author, query, id)

	if err != nil {
		return nil, err
	}

	return &author, nil
}
