package request

type PasswordReset struct {
	Email string `json:"email" form:"email" binding:"required"`
}

type NewPassword struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Code     string `json:"code" form:"code" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}
