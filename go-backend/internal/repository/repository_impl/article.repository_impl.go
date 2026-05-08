package repository_impl

import (
	"context"
	"go-backend/ent"
	"go-backend/internal/common/models"
	"go-backend/internal/common/pagination"
	"go-backend/internal/dto"
	"go-backend/internal/repository"

	"gorm.io/gorm"
)

type articleRepository struct {
	entClient  *ent.Client
	gormClient *gorm.DB
}

func NewArticleRepository(entClient *ent.Client, gormClient *gorm.DB) repository.ArticleRepository {
	return &articleRepository{
		entClient:  entClient,
		gormClient: gormClient,
	}
}

// Create implements [repository.ArticleRepository].
func (a *articleRepository) Create(ctx context.Context, body dto.ArticleCreateReq) (any, error) {
	entCreate := a.entClient.Articles.Create()

	entCreate = entCreate.SetTitle(body.Title)

	if body.Content != nil {
		entCreate = entCreate.SetContent(*body.Content)
	}

	if body.ImageUrl != nil {
		entCreate = entCreate.SetImageURL(*body.ImageUrl)
	}

	entCreate = entCreate.SetUserID(1)

	return entCreate.Save(ctx)

}

// CreateGorm implements [repository.ArticleRepository].
func (a *articleRepository) CreateGorm(ctx context.Context, body dto.ArticleCreateReq) (any, error) {
	article := models.Article{
		Title:    body.Title,
		Content:  body.Content,
		ImageUrl: body.ImageUrl,
	}

	result := gorm.WithResult()
	err := gorm.G[models.Article](a.gormClient, result).Create(ctx, &article)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetAll implements [repository.ArticleRepository].
func (a *articleRepository) GetAll(ctx context.Context, query pagination.Query) (any, error) {
	entQuery := a.entClient.Articles.Query()
	entQuery = entQuery.Limit(query.PageSize)
	entQuery = entQuery.Offset(query.Offset)
	return entQuery.All(ctx)
}
