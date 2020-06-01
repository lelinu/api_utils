package jwe

import (
	"github.com/lelinu/api_utils/utils/error_utils"
	"time"
)

//IService interface
type IService interface {
	GenerateJweToken(data map[string]interface{}) (string, *time.Time, *error_utils.ApiError)
	RefreshJweToken(token string) (string, *time.Time, *error_utils.ApiError)
}
