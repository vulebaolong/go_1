package middlewares

import (
	"errors"
	"fmt"
	"go-backend/internal/common/response"
	"go-backend/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	tokenUsecase usecase.TokenUsecase
}

func NewAuthMiddleware(tokenUsecase usecase.TokenUsecase) *AuthMiddleware {
	return &AuthMiddleware{
		tokenUsecase: tokenUsecase,
	}
}

func (a *AuthMiddleware) Protect(ctx *gin.Context) {
	accessToken, err := ctx.Cookie("accessToken")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			ctx.Error(response.NewUnauthorizedException())
			ctx.Abort()
			return
		}
		ctx.Error(response.NewBadRequestException())
		ctx.Abort()
		return
	}

	claim, err := a.tokenUsecase.VerifyAccessToken(accessToken)
	if err != nil {
		ctx.Error(response.NewUnauthorizedException())
		ctx.Abort()
		return
	}

	fmt.Println("accessToken", accessToken)
	fmt.Printf("claim: %+v", *claim)
}
