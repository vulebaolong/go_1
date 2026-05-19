package usecase_impl

import (
	"context"
	"go-backend/ent"
	"go-backend/internal/common/response"
	"go-backend/internal/dto"
	"go-backend/internal/repository"
	"go-backend/internal/usecase"

	"golang.org/x/crypto/bcrypt"
)

type authUsecase struct {
	userRepository repository.UserRepository
	tokenUsecase   usecase.TokenUsecase
}

func NewAuthUsecase(userRepository repository.UserRepository, tokenUsecase usecase.TokenUsecase) usecase.AuthUsecase {
	return &authUsecase{
		userRepository: userRepository,
		tokenUsecase:   tokenUsecase,
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

	// Mã hoá password
	// hash: băm (password) không dịch ngược
	// encryption: mã hoá chuyển dữ liệu
	hashPassByte, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}

	body.Password = string(hashPassByte)

	// tạo người dừng mới
	userNew, err := a.userRepository.CreateUser(ctx, body)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}

	return userNew, nil
}

// Login implements [usecase.AuthUsecase].
func (a *authUsecase) Login(ctx context.Context, body dto.AuthLoginReq) (*string, error) {
	user, err := a.userRepository.FindUserByEmail(ctx, body.Email)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}

	if user.Password == nil {
		return nil, response.NewBadRequestException("Vui lòng đăng nhập bằng google để cập nhật password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(body.Password))
	if err != nil {
		return nil, response.NewBadRequestException("Mật khẩu không chính xác")
	}

	// trả về accessToken và refreshToken
	token, err := a.tokenUsecase.CreateAccessToken(user.ID)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}

	return &token, nil
}

// GetInfo implements [usecase.AuthUsecase].
func (a *authUsecase) GetInfo(ctx context.Context) (*ent.Users, error) {
	return nil, nil
}
