package usecase

import "go-backend/internal/dto"

type TokenUsecase interface {
	CreateAccessToken(userId int) (string, error)
	VerifyAccessToken(accessToken string) (*dto.CustomClaim, error)
}
