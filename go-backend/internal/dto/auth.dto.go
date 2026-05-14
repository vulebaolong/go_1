package dto

type AuthRegisterReq struct {
	Email    string `json:"email" binding:"email,required"`
	FullName string `json:"fullName" binding:"required"`
	Password string `json:"password" binding:"required"`
}
