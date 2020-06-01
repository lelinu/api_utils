package firebase

import (
	"github.com/lelinu/golang-restclient/rest"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	longLink = "https://test.page.link?link=http://www.example.com/&apn=com.example.android&ibi=com.example.ios&data=akjflakjfladJdjdjdjdjj"
	baseUrl = "https://firebasedynamiclinks.googleapis.com/v1"
	url     = "/shortLinks?key=api_key"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestNewServiceInvalid(t *testing.T){
	service, err := NewService("")
	assert.NotNil(t, err)
	assert.Nil(t, service)
}

//TestShortenLinkInvalidKey
func TestShortenLinkInvalidKey(t *testing.T) {

	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:        baseUrl + url,
		HTTPMethod: http.MethodPost,
		RespBody: `{"error": {
					 "code": 400,
					 "message": "API key not valid. Please pass a valid API key.",
					 "status": "INVALID_ARGUMENT",
					 "details": [{
						"@type": "type.googleapis.com/google.rpc.Help",
						"links": [{
							"description": "Google developers console",
							"url": "https://console.developers.google.com"
						}]
					}]
				}
			}`,
		RespHTTPCode: http.StatusBadRequest,
	})

	service, err := NewService(baseUrl)
	assert.Nil(t, err)
	assert.NotNil(t, service)

	// assign values
	shortLink, err := service.ShortenLink(url, longLink)

	assert.NotNil(t, err)
	assert.EqualValues(t, "firebase: ShortenLink : API key not valid. Please pass a valid API key.", err.ErrorMessage)
	assert.EqualValues(t, "", shortLink)
}

//TestShortenLinkInvalidBadRequestResponse
func TestShortenLinkInvalidBadRequestResponse(t *testing.T) {

	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:        baseUrl + url,
		HTTPMethod: http.MethodPost,
		RespBody: `{"error": {
					 "status": 400,
					 "message": "API key not valid. Please pass a valid API key.",
					 "status": "INVALID_ARGUMENT",
					 "details": [{
						"@type": "type.googleapis.com/google.rpc.Help",
						"links": [{
							"description": "Google developers console",
							"url": "https://console.developers.google.com"
						}]
					}]
				}
			}`,
		RespHTTPCode: http.StatusBadRequest,
	})

	service, err := NewService(baseUrl)
	assert.Nil(t, err)
	assert.NotNil(t, service)

	// assign values
	shortLink, err := service.ShortenLink(url, longLink)

	assert.NotNil(t, err)
	assert.EqualValues(t, "firebase: ShortenLink : invalid rest error interface", err.ErrorMessage)
	assert.EqualValues(t, "", shortLink)
}

//TestShortenLinkInvalidOKResponse
func TestShortenLinkInvalidOKResponse(t *testing.T) {

	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:        baseUrl + url,
		HTTPMethod: http.MethodPost,
		RespBody: `{
			"ShortLink": true
		}`,
		RespHTTPCode: http.StatusOK,
	})

	service, err := NewService(baseUrl)
	assert.Nil(t, err)
	assert.NotNil(t, service)

	// assign values
	shortLink, err := service.ShortenLink(url, longLink)

	assert.NotNil(t, err)
	assert.EqualValues(t, "firebase: ShortenLink : failed to unmarshal response json: cannot unmarshal bool into Go struct field ShortLinksResponseModel.shortLink of type string", err.ErrorMessage)
	assert.EqualValues(t, "", shortLink)
}

//TestShortenLinkValid
func TestShortenLinkValid(t *testing.T) {

	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:        baseUrl + url,
		HTTPMethod: http.MethodPost,
		RespBody: `{
			"shortLink": "https://test.page.link/HyQM",
			"previewLink": "https://test.page.link/HyQM?d=1"
		}`,
		RespHTTPCode: http.StatusOK,
	})

	service, err := NewService(baseUrl)
	assert.Nil(t, err)
	assert.NotNil(t, service)

	// assign values
	shortLink, err := service.ShortenLink(url, longLink)

	assert.Nil(t, err)
	assert.NotNil(t, shortLink)
	assert.EqualValues(t, "https://test.page.link/HyQM", shortLink)
}