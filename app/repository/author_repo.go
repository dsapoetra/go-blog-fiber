package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"go-blog-fiber/app/model/repo"
)

type AuthorRepository struct {
	db *sqlx.DB
}

type IAuthorRepository interface {
	FindOneAuthor(id uuid.UUID) (*repo.Author, error)
	CreateAuthor(author *repo.Author) error
	FindAuthorByUsername(username string) (*repo.Author, error)
}

func NewAuthorRepository(db *sqlx.DB) IAuthorRepository {
	return &AuthorRepository{
		db: db,
	}
}

func (a *AuthorRepository) FindAuthorByUsername(username string) (*repo.Author, error) {
	author := repo.Author{}

	query := `SELECT * FROM authors where username = $1`

	err := a.db.Get(&author, query, username)

	if err != nil {
		return nil, err
	}

	return &author, nil
}

func (a *AuthorRepository) FindOneAuthor(id uuid.UUID) (*repo.Author, error) {
	author := repo.Author{}

	query := `SELECT * FROM authors where id = $1`

	err := a.db.Get(&author, query, id)

	if err != nil {
		return nil, err
	}

	return &author, nil
}

func (a *AuthorRepository) CreateAuthor(author *repo.Author) error {
	query := `INSERT INTO authors(created_at, updated_at, full_name, password, username, is_deleted) VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := a.db.Exec(query, author.CreatedAt, author.UpdatedAt, author.FullName, author.Password, author.UserName, false)
	if err != nil {
		fmt.Println("HEEEREE" + err.Error())
	}
	return err
}
