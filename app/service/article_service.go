package service

import (
	"github.com/google/uuid"
	"go-blog-fiber/app/model/repo"
	"go-blog-fiber/app/repository"
)

type ArticleService struct {
	db repository.IArticleRepository
}

type IArticleService interface {
	FindOneArticle(id uuid.UUID) (*repo.Article, error)
}

func NewArticleService(repo repository.IArticleRepository) IArticleService {
	return &ArticleService{
		db: repo,
	}
}

func (a *ArticleService) FindOneArticle(id uuid.UUID) (*repo.Article, error) {
	res, err := a.db.FindOneArticle(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}
