package smtp

import (
	"crypto/tls"
	"github.com/lelinu/api_utils/utils/error_utils"
	"gopkg.in/gomail.v2"
)

type Service struct {
	Host       string
	Port       int
	SenderName string
	Username   string
	Password   string
}

func NewService(host string, port int, senderName string, username string, password string) IService {
	return &Service{Host: host, Port: port, SenderName: senderName, Username: username, Password: password}
}

func (s *Service) Send(toList string, subject string, htmlBody string) *error_utils.ApiError {

	m := gomail.NewMessage()
	m.SetAddressHeader("From", s.Username, s.SenderName)
	m.SetHeader("To", toList)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", htmlBody)

	d := gomail.NewDialer(s.Host, s.Port, s.Username, s.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		return error_utils.NewInternalServerError(err.Error())
	}

	return nil
}
