package delivery

import (
	"go-backend/internal/handler"

	"github.com/gin-gonic/gin"
)

type articleDelivery struct {
	articleHandler *handler.ArticleHandler
}

func NewArticleDelivery(articleHandler *handler.ArticleHandler) *articleDelivery {
	return &articleDelivery{
		articleHandler: articleHandler,
	}
}

func (d *articleDelivery) RegisterRouter(apiGroup *gin.RouterGroup) {
	articleGroup := apiGroup.Group("article")
	{
		articleGroup.POST("", d.articleHandler.Create)
	}
}
