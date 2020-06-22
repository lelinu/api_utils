package email

import (
	"github.com/lelinu/api_utils/email/models"
	"github.com/lelinu/api_utils/utils/error_utils"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

//SendGridService struct
type SendGridService struct {
}

//NewSendGridService this method will return a new instance of PetService
func NewSendGridService() *SendGridService {
	return &SendGridService{}
}

//Send will send an email
func (s *SendGridService) Send(model *models.SendEmailModel) *error_utils.ApiError {

	from := mail.NewEmail(model.FromTag, model.From)
	subject := model.Subject
	toList := mail.NewEmail(model.ToList, model.ToList)

	m := mail.NewSingleEmail(from, subject, toList, model.PlainContent, model.HTMLContent)
	client := sendgrid.NewSendClient(model.Settings.Password)
	_, err := client.Send(m)
	if err != nil {
		return error_utils.NewInternalServerError(err.Error())
	}

	return nil
}
