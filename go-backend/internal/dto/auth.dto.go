package dto

type AuthRegisterReq struct {
	Email    string `json:"email" binding:"email,required"`
	FullName string `json:"fullName" binding:"required"`
	Password string `json:"password" binding:"required"`
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
