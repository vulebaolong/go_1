package delivery

import (
	"go-backend/internal/handler"

	"github.com/gin-gonic/gin"
)

type authDelivery struct {
	authHandler *handler.AuthHandler
}

func NewAuthDelivery(authHandler *handler.AuthHandler) *authDelivery {
	return &authDelivery{
		authHandler: authHandler,
	}
}

func (d *authDelivery) RegisterRouter(apiGroup *gin.RouterGroup) {
	authGroup := apiGroup.Group("auth")
	{
		authGroup.POST("register", d.authHandler.Register)
	}
}
