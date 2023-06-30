package request

type User struct {
	UserId     int64  `json:"userId,omitempty" form:"userId"`
	UserName   string `json:"userName" form:"userName" binding:"required"`
	Password   string `json:"password" form:"password" binding:"required"`
	MobileNo   string `json:"mobileNo,omitempty" form:"mobileNo"`
	Email      string `json:"email,omitempty" form:"email"`
	Gender     int    `json:"gender,omitempty" form:"gender"`
	UserAvatar string `json:"userAvatar,omitempty" form:"userAvatar"`
	IsVerified string `json:"isVerified,omitempty" form:"isVerified"`
}

type UserLogin struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Provider string `json:"provider"`
	MobileNo string `json:"mobileNo"`
}
