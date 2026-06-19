package delivery

import (
	"context"
	"go-order/internal/common/rabbitmq"
	"go-order/internal/handler"
)

type orderDelivery struct {
	orderHandler *handler.OrderHandler
}

func NewOrderDelivery(orderHandler *handler.OrderHandler) *orderDelivery {
	return &orderDelivery{
		orderHandler: orderHandler,
	}
}

func (d *orderDelivery) RegisterRouter(rabbitmq *rabbitmq.RabbitMQ) {
	rabbitmq.On(context.Background(), "CREATE_ORDER", d.orderHandler.Create)
}
