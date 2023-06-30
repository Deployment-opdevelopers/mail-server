package v1

import (
	"github.com/gin-gonic/gin"
)

func VerifyEmail(c *gin.Context) {
	// verificationFromContext, ok := c.Get("verification")
	// if !ok {
	// 	c.JSON(http.StatusBadRequest, util.SendErrorResponse(errors.New(constants.INVALID_REQUEST)))
	// 	return
	// }
	// verification := verificationFromContext.(*request.VerificationData)

	// if err := userServices.VerifyEmail(verification, verification.Type); err != nil {
	// 	c.JSON(http.StatusForbidden, util.SendErrorResponse(err))
	// 	return
	// }

	// res := response.Response{
	// 	Status:  constants.API_SUCCESS_STATUS,
	// 	Message: constants.USER_VERIFICATION_SUCCESS,
	// }
	// c.JSON(http.StatusOK, util.StructToJSON(res))
}
