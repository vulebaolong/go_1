package handler

import (
	"errors"
	"fmt"
	"go-backend/internal/common/helpers"
	"go-backend/internal/common/response"
	"go-backend/internal/dto"
	"go-backend/internal/usecase"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthHandler(authUsecase usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		authUsecase: authUsecase,
	}
}

func (a *AuthHandler) Register(ctx *gin.Context) {
	var body dto.AuthRegisterReq
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		if err == io.EOF {
			ctx.Error(response.NewBadRequestException("Body required"))
			return
		}
		ctx.Error(response.NewBadRequestException(err.Error()))
		return
	}

	result, err := a.authUsecase.Register(ctx.Request.Context(), body)
	if err != nil {
		ctx.Error(err)
		return
	}

	response.Success(result, "", 0, ctx)
}

func (a *AuthHandler) Login(ctx *gin.Context) {
	var body dto.AuthLoginReq
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		if err == io.EOF {
			ctx.Error(response.NewBadRequestException("Body required"))
			return
		}
		ctx.Error(response.NewBadRequestException(err.Error()))
		return
	}

	result, err := a.authUsecase.Login(ctx.Request.Context(), body)
	if err != nil {
		ctx.Error(err)
		return
	}

	setTokenCookie(ctx, result.AccessToken, result.RefreshToken)

	response.Success(true, "", 0, ctx)
}

func (a *AuthHandler) GetInfo(ctx *gin.Context) {
	user, err := helpers.GetUser(ctx)
	if err != nil {
		ctx.Error(response.NewBadRequestException(err.Error()))
		return
	}

	fmt.Println("helpers.GetUser", user)

	result, err := a.authUsecase.GetInfo(ctx.Request.Context(), user)
	if err != nil {
		ctx.Error(err)
		return
	}

	response.Success(result, "", 0, ctx)
}

func (a *AuthHandler) RefreshToken(ctx *gin.Context) {
	accessToken, err := ctx.Cookie("accessToken")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			ctx.Error(response.NewUnauthorizedException())
			return
		}
		ctx.Error(response.NewBadRequestException())
		return
	}
	refreshToken, err := ctx.Cookie("refreshToken")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			ctx.Error(response.NewUnauthorizedException())
			return
		}
		ctx.Error(response.NewBadRequestException())
		return
	}

	result, err := a.authUsecase.RefreshToken(ctx.Request.Context(), accessToken, refreshToken)
	if err != nil {
		ctx.Error(err)
		return
	}

	// Trường hợp 1: trả cả cặp token (accessToken, refreshToken)
	// kỹ thuận rotate
	// nếu người dùng không sử dụng trong vòng 24h, thì refreshToken không được làm mới => logout

	// Trường hợp 2: trả cả accessToken
	// từ khi người dùng login -> 24h sau -> logout

	setTokenCookie(ctx, result.AccessToken, result.RefreshToken)

	response.Success(result, "", 0, ctx)
}

func (a *AuthHandler) GoogleLogin(ctx *gin.Context) {
	result, err := a.authUsecase.GoogleLogin(ctx.Request.Context())
	if err != nil {
		ctx.Error(err)
		return
	}

	response.Success(result, "", 0, ctx)
}

func setTokenCookie(ctx *gin.Context, accessToken string, refreshToken string) {
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie(
		"accessToken",
		accessToken,
		0,
		"/",
		"",
		false,
		true,
	)
	ctx.SetCookie(
		"refreshToken",
		refreshToken,
		0,
		"/",
		"",
		false,
		true,
	)
}
