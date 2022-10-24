package service

import (
	"github.com/google/uuid"
	"go-blog-fiber/app/model"
	"go-blog-fiber/app/repository"
)

type AuthorService struct {
	db repository.IAuthorRepository
}

type IAuthorService interface {
	FindOneAuthor(id uuid.UUID) (*model.Author, error)
}

func NewAuthorService(repo repository.IAuthorRepository) IAuthorService {
	return &AuthorService{
		db: repo,
	}
}

func (a *AuthorService) FindOneAuthor(id uuid.UUID) (*model.Author, error) {
	res, err := a.db.FindOneAuthor(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}
