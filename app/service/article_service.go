package service

import (
	"github.com/google/uuid"
	"go-blog-fiber/app/model"
	"go-blog-fiber/app/repository"
	"log"
)

type ArticleService struct {
	db repository.IArticleRepository
}

type IArticleService interface {
	FindOneArticle(id uuid.UUID) (*model.Article, error)
}

func NewArticleService(repo repository.IArticleRepository) IArticleService {
	return &ArticleService{
		db: repo,
	}
}

func (a *ArticleService) FindOneArticle(id uuid.UUID) (*model.Article, error) {
	log.Println("HEREEEEEEEEEE 3")

	res, err := a.db.FindOneArticle(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}
