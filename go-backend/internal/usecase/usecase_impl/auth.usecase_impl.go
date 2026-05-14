package usecase_impl

import (
	"context"
	"go-backend/internal/common/response"
	"go-backend/internal/dto"
	"go-backend/internal/repository"
	"go-backend/internal/usecase"
)

type authUsecase struct {
	userRepository repository.UserRepository
}

func NewAuthUsecase(userRepository repository.UserRepository) usecase.AuthUsecase {
	return &authUsecase{
		userRepository: userRepository,
	}
}

// Register implements [usecase.AuthUsecase].
func (a *authUsecase) Register(ctx context.Context, body dto.AuthRegisterReq) (any, error) {
	// Kiểm tra user tồn tại hay chưa
	isExits, err := a.userRepository.ExitsByEmail(ctx, body.Email)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}

	// Nếu đã tồn tại thì trả lỗi
	if isExits {
		return nil, response.NewBadRequestException("Email đã có vui lòng đăng nhập")
	}

	return "Register", nil
}
