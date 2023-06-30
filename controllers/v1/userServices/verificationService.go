package userServices

import (
	"errors"
	"fmt"

	"opdevelopers.live/user/constants"
	"opdevelopers.live/user/helpers/request"
	"opdevelopers.live/user/services/discord"
	"opdevelopers.live/user/services/sendgrid"
)

func SendMail(req *request.VerificationData) error {
	from := "omprakash.bairwa.iitbhu20@gmail.com"
	to := []string{req.Email}
	mailType := sendgrid.MailConfirmation
	// otpCode := util.GenerateRandomString(6)
	var subject string
	var topMessage string
	var middleMessage string

	mailData := &sendgrid.MailData{
		Username:    req.UserName,
		LenderName:  req.LenderName,
		LenderEmail: req.LenderEmail,
		Company:     req.Company,
		Code:        req.Otp,
		Email:       "omprakash.bairwa.iitbhu20@gmail.com",
	}

	if req.Type == "otp" {
		mailType = sendgrid.PassWordReset
		subject = "Forgot Password"
	}
	if req.Type == "success" {
		mailType = sendgrid.MailConfirmation
		subject = "SignUp successfully"
	}
	if req.Type == "loan-accept" {
		mailType = 3
		subject = "Loan accepted"
		topMessage = fmt.Sprintf("Your loan request has been approved by %s", mailData.LenderName)
		middleMessage = fmt.Sprintf("Kindly check your profile and thanks for being a part of our %s organization.Lenders email id is %s", mailData.Company, mailData.LenderEmail)

	}
	if req.Type == "loan-reject" {
		mailType = 3
		subject = "Loan request rejected"
		// Sorry to say that your loan request has been rejected by {{lendersName}} So we can't proccess your loan application better luck next time and thanks for being a part of our {{company}} organization.
		topMessage = fmt.Sprintf("Sorry to say that your loan request has been rejected by %s", mailData.LenderName)
		middleMessage = fmt.Sprintf("So we can't proccess your loan application better luck next time and thanks for being a part of our %s organization.Lenders email id is %s", mailData.Company, mailData.LenderEmail)
	}
	if req.Type == "loan-modify" {
		mailType = 3
		subject = "Loan modify requested"
		// There is problem raised in your loan application kindly go through it . and update your application .
		topMessage = fmt.Sprintf("There is negotiation request raised by %s in your loan application.", mailData.LenderName)
		middleMessage = fmt.Sprintf("kindly go through it and update your application . Thanks for being a part of our %s organization.Lenders email id is %s", mailData.Company, mailData.LenderEmail)

	}
	if req.Type == "custom" {
		mailType = 3
		subject = req.Subject
	}

	mailReq := sendgrid.NewMail(from, to, subject, mailType, mailData)
	err := sendgrid.SendMail(mailReq, topMessage, middleMessage)
	if err != nil {
		discord.SendErrorMessageToDiscord(fmt.Sprintf("Error in sending verification mail to %s and the error is: %s", req.Email, err.Error()))
		return errors.New(constants.ERROR_IN_SENDING_MAIL)
	}
	return nil
}
