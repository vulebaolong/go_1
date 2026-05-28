package handler

import (
	"go-backend/internal/common/helpers"
	"go-backend/internal/common/response"
	"go-backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
	}
}

func (a *UserHandler) FindAll(ctx *gin.Context) {
	result, err := a.userUsecase.FindAll(ctx.Request.Context())
	if err != nil {
		ctx.Error(err)
		return
	}

	response.Success(result, "", 0, ctx)
}
func (a *UserHandler) AvatarLocal(ctx *gin.Context) {
	user, err := helpers.GetUser(ctx)
	if err != nil {
		ctx.Error(response.NewBadRequestException(err.Error()))
		return
	}

	fileHeader, err := ctx.FormFile("avatar")
	if err != nil {
		ctx.Error(response.NewBadRequestException(err.Error()))
		return
	}

	result, err := a.userUsecase.AvatarLocal(ctx.Request.Context(), fileHeader, user)
	if err != nil {
		ctx.Error(err)
		return
	}

	response.Success(result, "", 0, ctx)
}
func (a *UserHandler) AvatarCloud(ctx *gin.Context) {
	user, err := helpers.GetUser(ctx)
	if err != nil {
		ctx.Error(response.NewBadRequestException(err.Error()))
		return
	}

	fileHeader, err := ctx.FormFile("avatar")
	if err != nil {
		ctx.Error(response.NewBadRequestException(err.Error()))
		return
	}

	result, err := a.userUsecase.AvatarCloud(ctx.Request.Context(), fileHeader, user)
	if err != nil {
		ctx.Error(err)
		return
	}

	response.Success(result, "", 0, ctx)
}
