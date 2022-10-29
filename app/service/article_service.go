package service

import (
	"github.com/google/uuid"
	"go-blog-fiber/app/model/http"
	"go-blog-fiber/app/model/repo"
	"go-blog-fiber/app/repository"
	"time"
)

type ArticleService struct {
	db repository.IArticleRepository
}

type IArticleService interface {
	FindOneArticle(id uuid.UUID) (*repo.Article, error)
	CreateArticle(article *http.CreateArticleRequest) error
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

func (a *ArticleService) CreateArticle(articleRequest *http.CreateArticleRequest) error {
	article := &repo.Article{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Title:     articleRequest.Title,
		Content:   articleRequest.Content,
		Author:    articleRequest.Author,
	}

	err := a.db.CreateArticle(article)

	return err
}
