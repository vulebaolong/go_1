package repository_impl

import (
	"context"
	"go-order/ent"
	"go-order/internal/dto"
	"go-order/internal/repository"
)

type orderRepository struct {
	entClient *ent.Client
}

func NewOrderRepository(entClient *ent.Client) repository.OrderRepository {
	return &orderRepository{
		entClient: entClient,
	}
}

// FindAll implements [repository.OrderRepository].
func (a *orderRepository) FindAll(ctx context.Context) (any, error) {
	return nil, nil
}

// Create implements [repository.OrderRepository].
func (a *orderRepository) Create(ctx context.Context, body dto.CreateOrder) (any, error) {
	entCreate := a.entClient.Orders.Create()
	entCreate = entCreate.SetFoodID(body.FoodId)
	entCreate = entCreate.SetUserID(body.UserId)
	return entCreate.Save(ctx)

}
