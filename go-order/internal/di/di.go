package dependency

import (
	"go-order/ent"
	"go-order/internal/common/env"
	"go-order/internal/common/rabbitmq"
	"go-order/internal/delivery"
	"go-order/internal/handler"
	"go-order/internal/repository/repository_impl"
	"go-order/internal/usecase/usecase_impl"

	"github.com/gin-gonic/gin"
)

func Injection(ginEngine *gin.Engine, entClient *ent.Client, env *env.Env, allowOrigins []string, rabbitmq *rabbitmq.RabbitMQ) {

	orderRepository := repository_impl.NewOrderRepository(entClient)

	orderUsecase := usecase_impl.NewOrderUsecase(orderRepository)
	orderHandler := handler.NewOrderHandler(orderUsecase)
	orderDelivery := delivery.NewOrderDelivery(orderHandler)

	rootDelivery := delivery.NewRootDelivery(orderDelivery)
	rootDelivery.RegisterRouter(ginEngine, rabbitmq)
}
