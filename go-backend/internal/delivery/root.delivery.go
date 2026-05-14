package delivery

import (
	"github.com/gin-gonic/gin"
)

type rootDelivery struct {
	demoDelivery    *demoDelivery
	articleDelivery *articleDelivery
	authDelivery    *authDelivery
}

func NewRootDelivery(demoDelivery *demoDelivery, articleDelivery *articleDelivery, authDelivery *authDelivery) *rootDelivery {
	return &rootDelivery{
		demoDelivery:    demoDelivery,
		articleDelivery: articleDelivery,
		authDelivery:    authDelivery,
	}
}

func (r *rootDelivery) RegisterRouter(ginEngine *gin.Engine) {
	apiGroup := ginEngine.Group("api")
	{
		r.demoDelivery.RegisterRouter(apiGroup)
		r.articleDelivery.RegisterRouter(apiGroup)
		r.authDelivery.RegisterRouter(apiGroup)
		// gom các bộ API
	}
}
