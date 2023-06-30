package request

type FeedBack struct {
	Name    string `json:"name" form:"name" binding:"required"`
	Email   string `json:"email" form:"email" binding:"required"`
	Message string `json:"message" form:"message" binding:"required"`
}
