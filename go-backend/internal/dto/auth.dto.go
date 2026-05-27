package dto

type AuthRegisterReq struct {
	Email    string `json:"email" binding:"email,required"`
	FullName string `json:"fullName" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthCreateUserForGoogleReq struct {
	Email    string
	FullName string
	Avatar   string
	GoogleId string
}

type AuthLoginReq struct {
	Email    string `json:"email" binding:"email,required"`
	Password string `json:"password" binding:"required"`
}

type AuthLoginReturn struct {
	AccessToken  string
	RefreshToken string
}

type AuthRefreshTokenReturn struct {
	AccessToken  string
	RefreshToken string
}

type AuthGoogleLoginReturn struct {
	State string
	Url   string
}

type AuthGoogleCallbackInput struct {
	StateGoogle string
	StateCookie string
	Code        string
}
