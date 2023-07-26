package models

type ActionSignUpRequest struct {
	Email       string `json:"email" binding:"required,email"`
	Firstname   string `json:"firstName" binding:"required"`
	Lastname    string `json:"lastName" binding:"required"`
	DisplayName string `json:"displayName" binding:"required"`
}

type ActionSignInRequest struct {
	Email string `json:"email" binding:"required,email"`
}
