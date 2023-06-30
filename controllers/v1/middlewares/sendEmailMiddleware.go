package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"opdevelopers.live/user/constants"
	"opdevelopers.live/user/controllers/v1/utils"
)

func GetRequestBodyEmailData(verificationRequiredFields []string, verificationOptionalFields []string) gin.HandlerFunc {
	return func(context *gin.Context) {
		var requestObj interface{}

		// fmt.Println("level 1")
		if err := context.ShouldBind(&requestObj); err != nil {
			context.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
				"status":  constants.API_FAILED_STATUS,
				"message": constants.INVALID_REQUEST,
			})
			return
		}

		verificationJSON := requestObj.(map[string]interface{})
		// fmt.Println(verificationJSON)

		mailData, ok := utils.ValidateAndParseVerificationFields(verificationJSON, verificationRequiredFields, verificationOptionalFields)

		if !ok {
			context.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
				"status":  constants.API_FAILED_STATUS,
				"message": constants.INVALID_REQUEST,
			})
			return
		}

		context.Set("mailData", mailData)
		context.Next()
	}
}
