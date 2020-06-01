package firebase

import "github.com/lelinu/api_utils/utils/error_utils"

//IService interface
type IService interface {
	ShortenLink(url string, longLink string) (string, *error_utils.ApiError)
}
