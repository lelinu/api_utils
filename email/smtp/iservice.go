package smtp

import (
	"github.com/lelinu/api_utils/utils/error_utils"
)

type IService interface {
	Send(toList string, subject string, htmlBody string) *error_utils.ApiError
}
