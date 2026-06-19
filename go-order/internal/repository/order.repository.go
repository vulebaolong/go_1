package repository

import (
	"context"
	"go-order/internal/dto"
)

type OrderRepository interface {
	FindAll(ctx context.Context) (any, error)
	Create(ctx context.Context, body dto.CreateOrder) (any, error)
}
