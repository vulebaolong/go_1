package handler

import (
	"fmt"
	"go-backend/internal/common/response"
	"go-backend/internal/usecase"
	"path/filepath"
	"strings"
	"time"

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
	fileHeader, err := ctx.FormFile("avatar")

	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))

	fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)

	fmt.Println("Filename", fileHeader.Filename)
	fmt.Println("ext", ext)
	fmt.Println("fileName", fileName)

	fullPath := filepath.Join("public", "images", fileName)
	fmt.Println("fullPath", fullPath)

	ctx.SaveUploadedFile(fileHeader, fullPath)

	result, err := a.userUsecase.AvatarLocal(ctx.Request.Context())
	if err != nil {
		ctx.Error(err)
		return
	}

	response.Success(result, "", 0, ctx)
}
func (a *UserHandler) AvatarCloud(ctx *gin.Context) {
	result, err := a.userUsecase.AvatarCloud(ctx.Request.Context())
	if err != nil {
		ctx.Error(err)
		return
	}

	response.Success(result, "", 0, ctx)
}
