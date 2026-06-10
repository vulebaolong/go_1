package delivery

import (
	"github.com/gin-gonic/gin"
)

type rootDelivery struct {
	demoDelivery      *demoDelivery
	articleDelivery   *articleDelivery
	authDelivery      *authDelivery
	userDelivery      *userDelivery
	chatGroupDelivery *chatGroupDelivery
}

func NewRootDelivery(demoDelivery *demoDelivery, articleDelivery *articleDelivery, authDelivery *authDelivery, userDelivery *userDelivery, chatGroupDelivery *chatGroupDelivery) *rootDelivery {
	return &rootDelivery{
		demoDelivery:      demoDelivery,
		articleDelivery:   articleDelivery,
		authDelivery:      authDelivery,
		userDelivery:      userDelivery,
		chatGroupDelivery: chatGroupDelivery,
	}
}

func (r *rootDelivery) RegisterRouter(ginEngine *gin.Engine) {
	apiGroup := ginEngine.Group("api")
	{
		r.demoDelivery.RegisterRouter(apiGroup)
		r.articleDelivery.RegisterRouter(apiGroup)
		r.authDelivery.RegisterRouter(apiGroup)
		r.userDelivery.RegisterRouter(apiGroup)
		r.chatGroupDelivery.RegisterRouter(apiGroup)
		// gom các bộ API
	}
}
