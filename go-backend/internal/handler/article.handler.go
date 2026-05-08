package handler

import (
	"fmt"
	"go-backend/internal/common/pagination"
	"go-backend/internal/common/response"
	"go-backend/internal/dto"
	"go-backend/internal/usecase"
	"io"

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
	var body dto.ArticleCreateReq
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		if err == io.EOF {
			ctx.Error(response.NewBadRequestException("Body required"))
			return
		}
		ctx.Error(response.NewBadRequestException(err.Error()))
		return
	}

	reuslt, err := a.articleUsecase.Create(ctx, body)
	if err != nil {
		ctx.Error(err)
		return
	}

	response.Success(reuslt, "", 0, ctx)
}

func (a *ArticleHandler) FindAll(ctx *gin.Context) {
	queryPagi := pagination.Get(
		ctx.Query("page"),
		ctx.Query("pageSize"),
	)

	input := dto.ArticleFindAllInput{
		Query: queryPagi,
	}

	fmt.Printf("%+v \n\n", queryPagi)

	reuslt, err := a.articleUsecase.FindAll(ctx, input)
	if err != nil {
		ctx.Error(err)
		return
	}

	response.Success(reuslt, "", 0, ctx)
}
