package usecase_impl

import (
	"context"
	"go-backend/ent"
	"go-backend/internal/common/response"
	"go-backend/internal/dto"
	"go-backend/internal/repository"
	"go-backend/internal/usecase"

	"github.com/golang-jwt/jwt/v5"
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
func (a *authUsecase) Login(ctx context.Context, body dto.AuthLoginReq) (*dto.AuthLoginReturn, error) {
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
	accessToken, err := a.tokenUsecase.CreateAccessToken(user.ID)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}

	refreshToken, err := a.tokenUsecase.CreateRefreshToken(user.ID)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}

	return &dto.AuthLoginReturn{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// GetInfo implements [usecase.AuthUsecase].
func (a *authUsecase) GetInfo(ctx context.Context, user *ent.Users) (*ent.Users, error) {
	return user, nil
}

// RefreshToken implements [usecase.AuthUsecase].
func (a *authUsecase) RefreshToken(ctx context.Context, accessToken string, refreshToken string) (*dto.AuthRefreshTokenReturn, error) {
	// jwt.WithoutClaimsValidation(): không kiểm tra hết hạn
	claimAccessToken, err := a.tokenUsecase.VerifyAccessToken(accessToken, jwt.WithoutClaimsValidation())
	if err != nil {
		return nil, response.NewUnauthorizedException(err.Error())
	}

	claimRefreshToken, err := a.tokenUsecase.VerifyRefreshToken(refreshToken)
	if err != nil {
		return nil, response.NewUnauthorizedException(err.Error())
	}

	if claimAccessToken.UserId != claimRefreshToken.UserId {
		return nil, response.NewUnauthorizedException("2 Token không cùng 1 user")
	}

	user, err := a.userRepository.FindUserById(ctx, claimAccessToken.UserId)
	if err != nil {
		return nil, response.NewUnauthorizedException(err.Error())
	}

	accessTokenNew, err := a.tokenUsecase.CreateAccessToken(user.ID)
	if err != nil {
		return nil, response.NewUnauthorizedException(err.Error())
	}

	refreshTokenNew, err := a.tokenUsecase.CreateRefreshToken(user.ID)
	if err != nil {
		return nil, response.NewUnauthorizedException(err.Error())
	}

	return &dto.AuthRefreshTokenReturn{
		AccessToken:  accessTokenNew,
		RefreshToken: refreshTokenNew,
	}, nil
}

// GoogleLogin implements [usecase.AuthUsecase].
func (a *authUsecase) GoogleLogin(ctx context.Context) (any, error) {
	return "GoogleLogin", nil
}
