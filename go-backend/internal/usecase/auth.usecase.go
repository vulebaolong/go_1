package usecase

import (
	"context"
	"go-backend/ent"
	"go-backend/internal/dto"
)

type AuthUsecase interface {
	Register(ctx context.Context, body dto.AuthRegisterReq) (any, error)
	Login(ctx context.Context, body dto.AuthLoginReq) (*dto.AuthLoginReturn, error)
	GetInfo(ctx context.Context, user *ent.Users) (*ent.Users, error)
	RefreshToken(ctx context.Context, accessToken string, refreshToken string) (*dto.AuthRefreshTokenReturn, error)
}
