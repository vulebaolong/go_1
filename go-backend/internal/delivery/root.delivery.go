package delivery

import (
	"github.com/gin-gonic/gin"
)

type rootDelivery struct {
	demoDelivery        *demoDelivery
	articleDelivery     *articleDelivery
	authDelivery        *authDelivery
	userDelivery        *userDelivery
	chatGroupDelivery   *chatGroupDelivery
	chatMessageDelivery *chatMessageDelivery
}

func NewRootDelivery(demoDelivery *demoDelivery, articleDelivery *articleDelivery, authDelivery *authDelivery, userDelivery *userDelivery, chatGroupDelivery *chatGroupDelivery, chatMessageDelivery *chatMessageDelivery) *rootDelivery {
	return &rootDelivery{
		demoDelivery:        demoDelivery,
		articleDelivery:     articleDelivery,
		authDelivery:        authDelivery,
		userDelivery:        userDelivery,
		chatGroupDelivery:   chatGroupDelivery,
		chatMessageDelivery: chatMessageDelivery,
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
		r.chatMessageDelivery.RegisterRouter(apiGroup)
		// gom các bộ API
	}
}
