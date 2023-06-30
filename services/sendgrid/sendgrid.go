package sendgrid

import (
	"errors"
	"log"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"opdevelopers.live/user/config"
)

// CreateMail takes in a mail request and constructs a sendgrid mail type.
func CreateMail(mailReq *Mail, topMessage string, middleMessage string) []byte {
	config := config.Get()
	m := mail.NewV3Mail()
	from := mail.NewEmail(config.MailName, mailReq.From)
	m.SetFrom(from)

	if mailReq.Mtype == MailConfirmation {
		m.SetTemplateID(config.MailVerifictionTemplateID)
	} else if mailReq.Mtype == PassWordReset {
		m.SetTemplateID(config.PasswordResetTemplateID)
	} else if mailReq.Mtype == 3 {
		m.SetTemplateID(config.CustomTemplateID)

	}
	p := mail.NewPersonalization()

	tos := make([]*mail.Email, 0)
	for _, to := range mailReq.To {
		tos = append(tos, mail.NewEmail("user", to))
	}

	p.AddTos(tos...)

	p.SetDynamicTemplateData("code", mailReq.Data.Code)
	p.SetDynamicTemplateData("userName", mailReq.Data.Username)
	p.SetDynamicTemplateData("company", mailReq.Data.Company)
	p.SetDynamicTemplateData("email", mailReq.Data.Email)
	p.SetDynamicTemplateData("topMessage", topMessage)
	p.SetDynamicTemplateData("middleMessage", middleMessage)

	m.AddPersonalizations(p)
	return mail.GetRequestBody(m)
}

// SendMail creates a sendgrid mail from the given mail request and sends it.
func SendMail(mailReq *Mail, topMessage string, middleMessage string) error {
	config := config.Get()
	request := sendgrid.GetRequest(config.SendGridApiKey, "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	var Body = CreateMail(mailReq, topMessage, middleMessage)
	// log.Println(string(Body))
	request.Body = Body
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println("unable to send mail", "error", err)
		return err
	}

	if response.StatusCode >= 400 {
		log.Println("unable to send mail", "error", response.Body)
		return errors.New(response.Body)
	}

	log.Println("mail sent successfully", "sent status code", response.StatusCode, mailReq.To)
	return nil
}

// NewMail returns a new mail request.
func NewMail(from string, to []string, subject string, mailType MailType, data *MailData) *Mail {
	return &Mail{
		From:        from,
		To:          to,
		Subject:     subject,
		LenderName:  data.LenderName,
		LenderEmail: data.LenderName,
		Mtype:       mailType,
		Data:        data,
	}
}
