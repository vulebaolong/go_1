package usecase_impl

import (
	"context"
	"fmt"
	"go-order/internal/dto"
	"go-order/internal/repository"
	"go-order/internal/usecase"
)

type orderUsecase struct {
	orderRepository repository.OrderRepository
}

func NewOrderUsecase(orderRepository repository.OrderRepository) usecase.OrderUsecase {
	return &orderUsecase{
		orderRepository: orderRepository,
	}
}

// FindAll implements [usecase.OrderUsecase].
func (a *orderUsecase) FindAll(ctx context.Context) (any, error) {
	return "FindAll", nil
}

// Create implements [usecase.OrderUsecase].
func (a *orderUsecase) Create(ctx context.Context, body dto.CreateOrder) (any, error) {
	fmt.Println("body", body)
	return a.orderRepository.Create(ctx, body)
}
