package jwt

import (
	"github.com/lelinu/api_utils/utils/error_utils"
	"time"
)

//IService interface
type IService interface {
	GenerateJwtToken(customClaims map[string]interface{}) (string, *time.Time, *error_utils.ApiError)
	RefreshJwtToken(token string) (string, *time.Time, *error_utils.ApiError)
	ValidateJwtToken(token string) (map[string]interface{}, *error_utils.ApiError)
}
