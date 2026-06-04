package handler

import (
	"context"
	"go-backend/internal/usecase"
)

type ChatHandler struct {
	chatUsecase usecase.ChatUsecase
}

func NewChatHandler(chatUsecase usecase.ChatUsecase) *ChatHandler {
	return &ChatHandler{
		chatUsecase: chatUsecase,
	}
}

func (c *ChatHandler) CreateGroup() {
	c.chatUsecase.CreateGroup(context.Background())
}
