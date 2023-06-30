package request

type GetInTouch struct {
	UserId int64  `json:"userId" form:"userId" binding:"required"`
	Email  string `json:"email" form:"email" binding:"required"`
}
