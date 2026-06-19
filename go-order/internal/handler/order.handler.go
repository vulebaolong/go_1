package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"go-order/internal/common/response"
	"go-order/internal/dto"
	"go-order/internal/usecase"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderUsecase usecase.OrderUsecase
}

func NewOrderHandler(orderUsecase usecase.OrderUsecase) *OrderHandler {
	return &OrderHandler{
		orderUsecase: orderUsecase,
	}
}

func (a *OrderHandler) FindAll(ctx *gin.Context) {
	result, err := a.orderUsecase.FindAll(ctx.Request.Context())
	if err != nil {
		ctx.Error(err)
		return
	}

	response.Success(result, "", 0, ctx)
}

func (a *OrderHandler) Create(ctx context.Context, payload []byte) error {
	var body dto.CreateOrder
	err := json.Unmarshal(payload, &body)
	if err != nil {
		return err
	}

	_, err = a.orderUsecase.Create(ctx, body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
