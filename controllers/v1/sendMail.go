package v1

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"opdevelopers.live/user/constants"
	"opdevelopers.live/user/controllers/v1/userServices"
	"opdevelopers.live/user/helpers/request"
	"opdevelopers.live/user/helpers/response"
	"opdevelopers.live/user/helpers/util"
)

func SendEmailToUser(c *gin.Context) {
	mailDataFromContext, ok := c.Get("mailData")
	// fmt.Println(mailDataFromContext)
	if !ok {
		c.JSON(http.StatusBadRequest, util.SendErrorResponse(errors.New(constants.INVALID_REQUEST)))
		return
	}
	mailData := mailDataFromContext.(*request.VerificationData)
	admin := util.IsAdmin(mailData)

	if !admin {
		c.JSON(http.StatusUnauthorized, util.SendErrorResponse(errors.New(constants.NOT_AUTHERIZED)))
		return
	}
	// var otpCode int
	err := userServices.SendMail(mailData)

	if err != nil {
		c.JSON(http.StatusForbidden, util.SendErrorResponse(err))
		return
	}

	res := response.Response{
		Status:  constants.API_SUCCESS_STATUS,
		Message: constants.SEND_EMAIL_SUCCESS,
	}
	c.JSON(http.StatusOK, util.StructToJSON(res))

}
