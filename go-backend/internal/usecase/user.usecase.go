package usecase

import (
	"context"
	"go-backend/ent"
	"mime/multipart"
)

type UserUsecase interface {
	FindAll(ctx context.Context) (any, error)
	AvatarLocal(ctx context.Context, fileHeader *multipart.FileHeader, user *ent.Users) (any, error)
	AvatarCloud(ctx context.Context, fileHeader *multipart.FileHeader, user *ent.Users) (any, error)
}
