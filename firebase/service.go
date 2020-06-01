package firebase

import (
	"encoding/json"
	"fmt"
	"github.com/lelinu/api_utils/firebase/models"
	"github.com/lelinu/api_utils/utils/error_utils"
	"github.com/lelinu/golang-restclient/rest"
	"strings"
	"time"
)



//Service object for injection
type Service struct {
	restClient rest.RequestBuilder
}

//NewService this method will return a new instance of ParameterStoreService
func NewService(baseUrl string) (*Service, *error_utils.ApiError) {
	var service = &Service{}
	err := service.init(baseUrl)
	if err != nil{
		return nil, err
	}

	return service, err
}

//init will initialize defaults
func (a *Service) init(baseUrl string) *error_utils.ApiError {

	if len(strings.TrimSpace(baseUrl)) == 0 {
		return error_utils.NewBadRequestError("Firebase: base URL is empty")
	}

	// init rest client
	var restClient = rest.RequestBuilder{
		Timeout: 100 * time.Millisecond,
		BaseURL: baseUrl,
	}

	a.restClient = restClient

	return nil
}

//ShortenLink this method will be used to shorten the link
func (a *Service) ShortenLink(url string, longLink string) (string, *error_utils.ApiError) {

	// Creating body json values
	request := &models.ShortLinksRequestModel{
		LongDynamicLink: longLink,
		Suffix: models.ShortLinksOptionRequestModel{
			Option: "SHORT",
		},
	}
	// Marshall request to json
	requestJSONValues, _ := json.Marshal(request)
	resp := a.restClient.Post(url, requestJSONValues)

	// Check response code
	if resp.StatusCode > 299 {
		var firebaseError models.ShortLinksErrorResponseModel
		if err := json.Unmarshal(resp.Bytes(), &firebaseError); err != nil {
			return "", error_utils.NewInternalServerError("firebase: ShortenLink : invalid rest error interface")
		}
		return "", error_utils.NewInternalServerError(fmt.Sprintf("firebase: ShortenLink : %v", firebaseError.Error.Message))
	}

	// Bind to response
	var firebaseResponse models.ShortLinksResponseModel
	err := json.Unmarshal(resp.Bytes(), &firebaseResponse)
	if err != nil {
		return  "", error_utils.NewInternalServerError(fmt.Sprintf("firebase: ShortenLink : failed to unmarshal response %v", err))
	}

	return firebaseResponse.ShortLink, nil
}
