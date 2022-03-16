package models

type LoginRequest struct {
	Username   string `json:"UserName" form:"UserName" binding:"required"`
	Password   string `json:"Password" form:"Password" binding:"required"`
	RememberMe bool   `json:"RememberMe" form:"RememberMe"`
}
