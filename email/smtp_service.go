package email

import (
	"crypto/tls"
	"github.com/lelinu/api_utils/email/models"
	"github.com/lelinu/api_utils/utils/error_utils"
	"gopkg.in/gomail.v2"
)

//STMPService struct
type STMPService struct {
}

//NewSMTPService this method will return a new instance of PetService
func NewSMTPService() *STMPService {
	return &STMPService{}
}

//Send will send an email with html body
func (s *STMPService) Send(model *models.SendEmailModel) *error_utils.ApiError {

	m := gomail.NewMessage()
	m.SetHeader("From", model.From)
	m.SetHeader("To",  model.ToList)
	m.SetHeader("Subject", model.Subject)
	m.SetBody("text/html", model.HTMLContent)

	d := gomail.NewDialer(model.Settings.Host, model.Settings.Port, model.Settings.Username, model.Settings.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		return error_utils.NewInternalServerError(err.Error())
	}

	return nil
}
