package sendgrid

import (
	"github.com/lelinu/api_utils/utils/error_utils"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

//Service struct
type Service struct {
	ApiKey string
	FromTag string
	From string
}

func NewService(apiKey string, fromTag string, from string) IService {
	return &Service{ApiKey: apiKey, FromTag: fromTag, From: from}
}

func (s *Service) Send(toList string, subject string, htmlBody string, plainBody string) *error_utils.ApiError {

	from := mail.NewEmail(s.FromTag, s.From)
	to := mail.NewEmail(toList, toList)

	m := mail.NewSingleEmail(from, subject, to, plainBody, htmlBody)
	client := sendgrid.NewSendClient(s.ApiKey)
	_, err := client.Send(m)
	if err != nil {
		return error_utils.NewInternalServerError(err.Error())
	}

	return nil
}
