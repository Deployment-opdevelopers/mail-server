package request

type Profile struct {
	UserId   int64  `json:"userId,omitempty" form:"userId"`
	UserName string `json:"userName" form:"userName" binding:"required"`
	Gender   int    `json:"gender" form:"gender" binding:"required"`
}
