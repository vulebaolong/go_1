package usecase_impl

import (
	"context"
	"go-backend/internal/repository"
	"go-backend/internal/usecase"
)

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) usecase.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

// FindAll implements [usecase.UserUsecase].
func (a *userUsecase) FindAll(ctx context.Context) (any, error) {
	return "FindAll", nil
}

// AvatarLocal implements [usecase.UserUsecase].
func (a *userUsecase) AvatarLocal(ctx context.Context) (any, error) {
	return "AvatarLocal", nil
}

// AvatarCloud implements [usecase.UserUsecase].
func (a *userUsecase) AvatarCloud(ctx context.Context) (any, error) {
	return "AvatarCloud", nil
}
