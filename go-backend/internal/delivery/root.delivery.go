package delivery

import (
	"github.com/gin-gonic/gin"
)

type rootDelivery struct {
	demoDelivery    *demoDelivery
	articleDelivery *articleDelivery
}

func NewRootDelivery(demoDelivery *demoDelivery, articleDelivery *articleDelivery) *rootDelivery {
	return &rootDelivery{
		demoDelivery:    demoDelivery,
		articleDelivery: articleDelivery,
	}
}

func (r *rootDelivery) RegisterRouter(ginEngine *gin.Engine) {
	apiGroup := ginEngine.Group("api")
	{
		r.demoDelivery.RegisterRouter(apiGroup)
		r.articleDelivery.RegisterRouter(apiGroup)
		// gom các bộ API
	}
}
