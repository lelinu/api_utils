package email

import "github.com/lelinu/api_utils/email/models"

//IService interface
type IService interface {
	Send(model *models.SendEmailModel) error
}
