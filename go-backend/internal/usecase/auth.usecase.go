package usecase

import (
	"context"
	"go-backend/ent"
	"go-backend/internal/dto"
)

type AuthUsecase interface {
	Register(ctx context.Context, body dto.AuthRegisterReq) (any, error)
	Login(ctx context.Context, body dto.AuthLoginReq) (*string, error)
	GetInfo(ctx context.Context) (*ent.Users, error)
}
