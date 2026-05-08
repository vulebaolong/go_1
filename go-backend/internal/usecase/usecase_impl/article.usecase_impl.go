package usecase_impl

import (
	"context"
	"go-backend/internal/common/response"
	"go-backend/internal/dto"
	"go-backend/internal/repository"
	"go-backend/internal/usecase"
)

type articleUsecase struct {
	articleRepository repository.ArticleRepository
}

func NewArticleUsecase(articleRepository repository.ArticleRepository) usecase.ArticleUsecase {
	return &articleUsecase{
		articleRepository: articleRepository,
	}
}

// Create implements [usecase.ArticleUsecase].
func (a *articleUsecase) Create(ctx context.Context, body dto.ArticleCreateReq) (any, error) {
	data, err := a.articleRepository.Create(ctx, body)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}
	return data, nil
}

// FindAll implements [usecase.ArticleUsecase].
func (a *articleUsecase) FindAll(ctx context.Context, input dto.ArticleFindAllInput) (any, error) {
	data, err := a.articleRepository.GetAll(ctx, input.Query)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}
	return data, nil
}
