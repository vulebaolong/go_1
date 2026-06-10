package usecase

import (
	"context"
	"go-backend/ent"
)

type ChatUsecase interface {
	CreateGroup(ctx context.Context, accessToken string, targetUserIds []int, name string) (*ent.ChatGroups, error)
	JoinGroup(ctx context.Context, accessToken string, chatGroupId int) (*ent.ChatGroups, error)
}
