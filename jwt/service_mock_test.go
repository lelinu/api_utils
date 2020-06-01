package jwt

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	funcGenerateJwtToken func(data map[string]interface{}) (string, time.Time, error)
	funcRefreshJwtToken  func(token string) (string, time.Time, error)
)

type authServiceMock struct{}

func (a *authServiceMock) GenerateJwtToken(data map[string]interface{}) (string, time.Time, error) {
	return funcGenerateJwtToken(data)
}

func (a *authServiceMock) RefreshJwtToken(token string) (string, time.Time, error) {
	return funcRefreshJwtToken(token)
}

//TestGenerateJwtTokenMockValid will test the Generate Jwt Token method
//go test -v
func TestGenerateJwtTokenMockValid(t *testing.T) {

	service := &authServiceMock{}
	// return valid parameters
	funcGenerateJwtToken = func(data map[string]interface{}) (string, time.Time, error) {
		return "This is a jwt token", time.Now(), nil
	}

	token, expires, err := service.GenerateJwtToken(map[string]interface{}{})

	assert.NotNil(t, token)
	assert.NotNil(t, expires)
	assert.Nil(t, err)
	assert.EqualValues(t, "This is a jwt token", token)
}

//TestRefreshJwtTokenMockValid will test the Refresh Jwt Token method
//go test -v
func TestRefreshJwtTokenMockValid(t *testing.T) {
	service := &authServiceMock{}

	// return valid parameters
	funcRefreshJwtToken = func(token string) (string, time.Time, error) {
		return "This is a refreshed jwt token", time.Now(), nil
	}

	token, expires, err := service.RefreshJwtToken("This is a jwt token")

	assert.NotNil(t, token)
	assert.NotNil(t, expires)
	assert.Nil(t, err)
	assert.EqualValues(t, "This is a refreshed jwt token", token)
}
