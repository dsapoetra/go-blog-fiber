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
	CreateAuthor(author *model.Author) error
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

func (a *AuthorRepository) CreateAuthor(author *model.Author) error {
	query := `INSERT INTO authors(created_at, updated_at, full_name) VALUES ($1, $2, $3)`

	_, err := a.db.Exec(query, author.CreatedAt, author.UpdatedAt, author.FullName)

	return err
}
