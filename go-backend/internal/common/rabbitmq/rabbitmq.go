package rabbitmq

import (
	"fmt"
	"go-backend/internal/common/env"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	Conn *amqp.Connection
}

// doc https://www.rabbitmq.com/tutorials/tutorial-six-go
// mẫu code cho client: https://github.com/rabbitmq/rabbitmq-tutorials/blob/main/go/rpc_client.go
func NewRabbitMQ(env *env.Env) *RabbitMQ {
	conn, err := amqp.Dial(env.RabbitMQURL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("✅ [RABBIT_MQ] Connection To RabbitMQ Successfully")

	return &RabbitMQ{
		Conn: conn,
	}
}
