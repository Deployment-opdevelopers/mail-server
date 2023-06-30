package request

type Admin struct {
	Email string `json:"email" form:"email" binding:"required"`
}