package usecase_impl

import (
	"context"
	"go-backend/internal/usecase"
)

type ChatUsecase struct{}

func NewChatUsecase() usecase.ChatUsecase {
	return &ChatUsecase{}
}

// CreateGroup implements [usecase.ChatUsecase].
func (c *ChatUsecase) CreateGroup(ctx context.Context) (any, error) {
	panic("unimplemented")
}
