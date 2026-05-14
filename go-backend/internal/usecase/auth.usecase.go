package usecase

import (
	"context"
	"go-backend/internal/dto"
)

type AuthUsecase interface {
	Register(ctx context.Context, body dto.AuthRegisterReq) (any, error)
}
