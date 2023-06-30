package response

type ProfileDetails struct {
	UserId   int64  `json:"userId,omitempty" form:"userId"`
	UserName string `json:"userName" form:"userName" binding:"required"`
	Email    string `json:"email,omitempty" form:"email"`
	MobileNo string `json:"mobileNo,omitempty" form:"mobileNo"`
	Gender   int    `json:"gender,omitempty" form:"gender"`
}
