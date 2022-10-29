package service

import (
	"github.com/google/uuid"
	"go-blog-fiber/app/model/http"
	"go-blog-fiber/app/model/repo"
	"go-blog-fiber/app/repository"
	"go-blog-fiber/pkg/utils"
)

type AuthorService struct {
	db repository.IAuthorRepository
}

type IAuthorService interface {
	FindOneAuthor(id uuid.UUID) (*repo.Author, error)
	CreateAuthor(author *repo.Author) error
	LoginAuthor(credential *http.LoginAuthorRequest) (string, error)
}

func NewAuthorService(repo repository.IAuthorRepository) IAuthorService {
	return &AuthorService{
		db: repo,
	}
}

func (a *AuthorService) LoginAuthor(credential *http.LoginAuthorRequest) (string, error) {
	author, err := a.db.FindAuthorByUsername(credential.UserName)

	if err != nil {
		return "", err
	}

	if err = utils.ComparePass(author.Password, credential.Password); err != nil {
		return "", err
	}

	tokenString, err := utils.SignedToken(*author)

	if err != nil {
		return "", err
	}

	return tokenString, err
}

func (a *AuthorService) FindOneAuthor(id uuid.UUID) (*repo.Author, error) {
	res, err := a.db.FindOneAuthor(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *AuthorService) CreateAuthor(author *repo.Author) error {
	err := a.db.CreateAuthor(author)

	return err
}
