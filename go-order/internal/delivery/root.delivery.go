package delivery

import (
	"go-order/internal/common/rabbitmq"

	"github.com/gin-gonic/gin"
)

type rootDelivery struct {
	orderDelivery *orderDelivery
}

func NewRootDelivery(orderDelivery *orderDelivery) *rootDelivery {
	return &rootDelivery{
		orderDelivery: orderDelivery,
	}
}

func (r *rootDelivery) RegisterRouter(ginEngine *gin.Engine, rabbitmq *rabbitmq.RabbitMQ) {
	r.orderDelivery.RegisterRouter(rabbitmq)
}
