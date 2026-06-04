package usecase

import "context"

type ChatUsecase interface {
	CreateGroup(ctx context.Context) (any, error)
}
