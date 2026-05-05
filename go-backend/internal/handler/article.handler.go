package handler

import (
	"go-backend/internal/common/response"
	"go-backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

type ArticleHandler struct {
	articleUsecase usecase.ArticleUsecase
}

func NewArticleHandler(articleUsecase usecase.ArticleUsecase) *ArticleHandler {
	return &ArticleHandler{
		articleUsecase: articleUsecase,
	}
}

func (a *ArticleHandler) Create(ctx *gin.Context) {
	reuslt, err := a.articleUsecase.Create()
	if err != nil {
		ctx.Error(err)
		return
	}

	response.Success(reuslt, "", 0, ctx)
}
