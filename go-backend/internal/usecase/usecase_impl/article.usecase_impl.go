package usecase_impl

import "go-backend/internal/usecase"

type articleUsecase struct{}

func NewArticleUsecase() usecase.ArticleUsecase {
	return &articleUsecase{}
}

// Create implements [usecase.ArticleUsecase].
func (a *articleUsecase) Create() (any, error) {
	return "Create", nil
}
