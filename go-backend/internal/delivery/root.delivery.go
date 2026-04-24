package delivery

import (
	"github.com/gin-gonic/gin"
)

type rootDelivery struct {
	demoDelivery *demoDelivery
}

func NewRootDelivery(demoDelivery *demoDelivery) *rootDelivery {
	return &rootDelivery{
		demoDelivery: demoDelivery,
	}
}

func (r *rootDelivery) RegisterRouter(ginEngine *gin.Engine) {
	apiGroup := ginEngine.Group("api")
	{

		r.demoDelivery.RegisterRouter(apiGroup)
		// gom các bộ API
	}
}
