package usecase

import (
	"context"
)

type UserUsecase interface {
	FindAll(ctx context.Context) (any, error)
	AvatarLocal(ctx context.Context) (any, error)
	AvatarCloud(ctx context.Context) (any, error)
}
