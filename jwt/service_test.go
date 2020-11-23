package jwt

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	signingAlgorithm string
	jwtSecretKey     string
	issuer           string
	customClaims     map[string]interface{}
)

func init() {
	signingAlgorithm = "HS256"
	jwtSecretKey = "s4IIq9lQm2SKBlJoHAWzkRGSNaPCLZw2Ed927XEcBMrvqyU0wpPgTttj2HAvYb9S"
	issuer = "lelinu"
	customClaims = map[string]interface{}{"id": 1, "role": "user", "merchantID": "1234546"}
}

func TestNewServiceInvalidAlgorithm(t *testing.T) {

	_, err := NewService("invalid-algo", jwtSecretKey, issuer, 1, 1)
	assert.NotNil(t, err)
	assert.EqualValues(t, "Auth: invalid signing algorithm", err.ErrorMessage)
}

//TestInitWithInvalidAlgorithm
func TestInitWithInvalidAlgorithm(t *testing.T) {

	service := Service{}
	err := service.init("invalid-algo", jwtSecretKey, issuer, 1, 1)

	assert.NotNil(t, err)
	assert.EqualValues(t, "Auth: invalid signing algorithm", err.ErrorMessage)
}

//TestInitWithInvalidIssuer
func TestInitWithInvalidIssuer(t *testing.T) {

	service := Service{}
	err := service.init(signingAlgorithm, jwtSecretKey, "", 1, 1)

	assert.NotNil(t, err)
	assert.EqualValues(t, "Auth: issuer cannot be empty", err.ErrorMessage)
}

//TestInitWithInvalidMaxRefresh
func TestInitWithInvalidMaxRefresh(t *testing.T) {

	service := Service{}
	err := service.init(signingAlgorithm, jwtSecretKey, issuer, 1, 0)

	assert.NotNil(t, err)
	assert.EqualValues(t, "Auth: max refresh should be greater than 0", err.ErrorMessage)
}

//TestInitWithInvalidSecretKey
func TestInitWithInvalidSecretKey(t *testing.T) {

	service := Service{}
	err := service.init(signingAlgorithm, " ", issuer, 1, 1)

	assert.NotNil(t, err)
	assert.EqualValues(t, "Auth: jwt Secret key cannot be empty", err.ErrorMessage)
}

//TestGenerateJwtTokenValid
func TestGenerateJwtTokenValidWithCustomClaims(t *testing.T) {

	service, err := NewService(signingAlgorithm, jwtSecretKey, issuer, 1, 1)
	assert.Nil(t, err)
	assert.NotNil(t, service)

	token, time, err := service.GenerateJwtToken(customClaims)

	assert.NotNil(t, token)
	assert.NotNil(t, time)
	assert.Nil(t, err)
}

//TestGenerateJwtTokenValidWithoutCustomClaims
func TestGenerateJwtTokenValidWithoutCustomClaims(t *testing.T) {

	// arrange
	service, err := NewService(signingAlgorithm, jwtSecretKey, issuer, 1, 1)
	assert.Nil(t, err)
	assert.NotNil(t, service)

	// act
	token, time, err := service.GenerateJwtToken(nil)

	// assert
	assert.NotNil(t, token)
	assert.NotNil(t, time)
	assert.Nil(t, err)
}

//TestRefreshJwtTokenValid
func TestRefreshJwtTokenValidWithCustomClaims(t *testing.T) {

	service, err := NewService(signingAlgorithm, jwtSecretKey, issuer, 1, 1)
	token, time, err := service.GenerateJwtToken(customClaims)

	assert.NotNil(t, token)
	assert.NotNil(t, time)
	assert.Nil(t, err)

	newToken, newTime, newErr := service.RefreshJwtToken(token)

	assert.NotNil(t, newToken)
	assert.NotNil(t, newTime)
	assert.Nil(t, newErr)
}

//TestRefreshJwtTokenValidWithoutCustomClaims
func TestRefreshJwtTokenValidWithoutCustomClaims(t *testing.T) {

	service, err := NewService(signingAlgorithm, jwtSecretKey, issuer, 1, 1)
	token, time, err := service.GenerateJwtToken(nil)

	assert.NotNil(t, token)
	assert.NotNil(t, time)
	assert.Nil(t, err)

	newToken, newTime, newErr := service.RefreshJwtToken(token)

	assert.NotNil(t, newToken)
	assert.NotNil(t, newTime)
	assert.Nil(t, newErr)
}

//TestValidateJwtTokenInvalidIssuer
func TestValidateJwtTokenInvalidIssuer(t *testing.T) {
	// generate with lelinu as issuer
	gService, err := NewService(signingAlgorithm, jwtSecretKey, issuer, 1, 1)
	assert.Nil(t, err)

	token, _, err := gService.GenerateJwtToken(customClaims)
	assert.Nil(t, err)
	assert.NotEqual(t, "", token)

	// validate with test as issuer
	vService, err := NewService(signingAlgorithm, jwtSecretKey, "test", 1, 1)
	assert.Nil(t, err)

	claims, err := vService.ValidateJwtToken(token)

	assert.NotNil(t, err)
	assert.EqualValues(t, "Invalid issuer", err.ErrorMessage)
	assert.Nil(t, claims)
}

//TestValidateJweTokenExpiredWithoutCustomClaims
func TestValidateJwtTokenExpiredWithoutCustomClaims(t *testing.T) {

	// arrange
	expiredToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTA2ODQ3MDcsImlzcyI6ImxlbGludSIsIm9yaWdfaWF0IjoxNTkwNjg0NjQ3fQ.CZ2zg17YDfv6ocG-L6PjF3IjxGjE5NnihsjJs2X6Eaw"
	service, err := NewService(signingAlgorithm, jwtSecretKey, issuer, 1, 1)
	assert.Nil(t, err)
	assert.NotNil(t, service)

	// act
	claims, err := service.ValidateJwtToken(expiredToken)

	// assert
	assert.NotNil(t, err)
	assert.EqualValues(t, "Token is expired", err.ErrorMessage)
	assert.Nil(t, claims)
}

//TestValidateJwtTokenExpiredWithCustomClaimsf
func TestValidateJwtTokenExpiredWithCustomClaims(t *testing.T) {

	// arrange
	expiredToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTA2ODQ2NTMsImlkIjoxLCJpc3MiOiJsZWxpbnUiLCJtZXJjaGFudElEIjoiMTIzNDU0NiIsIm9yaWdfaWF0IjoxNTkwNjg0NTkzLCJyb2xlIjoidXNlciJ9.5YjIqj3479ahgOknuY_9B9zkIP3FzOVq-r6eAiBGw04"
	service, err := NewService(signingAlgorithm, jwtSecretKey, issuer, 1, 1)
	assert.Nil(t, err)
	assert.NotNil(t, service)

	// act
	claims, err := service.ValidateJwtToken(expiredToken)

	// assert
	assert.NotNil(t, err)
	assert.EqualValues(t, "Token is expired", err.ErrorMessage)
	assert.Nil(t, claims)
}

//TestValidateJwtTokenValidWithCustomClaims
func TestValidateJwtTokenValidWithCustomClaims(t *testing.T) {

	service, err := NewService(signingAlgorithm, jwtSecretKey, issuer, 1, 1)
	assert.Nil(t, err)
	assert.NotNil(t, service)

	// act
	token, exp, err := service.GenerateJwtToken(customClaims)
	assert.NotNil(t, token)
	assert.NotNil(t, exp)
	assert.Greater(t, exp.Unix(), time.Now().UTC().Unix())
	assert.Nil(t, err)

	// act
	claims, err := service.ValidateJwtToken(token)

	// assert
	assert.Nil(t, err)
	assert.NotNil(t, claims)
	assert.NotNil(t, claims["exp"])
	assert.NotNil(t, claims["id"])
	assert.NotNil(t, claims["iss"])
	assert.NotNil(t, claims["merchantID"])
	assert.NotNil(t, claims["orig_iat"])
	assert.NotNil(t, claims["role"])
}

//TestValidateJwtTokenValidWithoutCustomClaims
func TestValidateJwtTokenValidWithoutCustomClaims(t *testing.T) {

	service, err := NewService(signingAlgorithm, jwtSecretKey, issuer, 1, 1)
	assert.Nil(t, err)
	assert.NotNil(t, service)

	// act
	token, exp, err := service.GenerateJwtToken(nil)
	assert.NotNil(t, token)
	assert.NotNil(t, exp)
	assert.Greater(t, exp.Unix(), time.Now().UTC().Unix())
	assert.Nil(t, err)

	// act
	claims, err := service.ValidateJwtToken(token)

	// assert
	assert.Nil(t, err)
	assert.NotNil(t, claims)
	assert.NotNil(t, claims["exp"])
	assert.Nil(t, claims["id"])
	assert.NotNil(t, claims["iss"])
	assert.Nil(t, claims["merchantID"])
	assert.NotNil(t, claims["orig_iat"])
	assert.Nil(t, claims["role"])
}

//TestParseTokenStringInvalid
func TestParseTokenStringInvalid(t *testing.T) {

	token := "invalid-token"
	service, apiErr := NewService(signingAlgorithm, jwtSecretKey, issuer, 1, 1)
	assert.Nil(t, apiErr)
	assert.NotNil(t, service)

	jwtToken, err := service.parseTokenString(token)
	assert.NotNil(t, err)
	assert.EqualValues(t, "token contains an invalid number of segments", err.Error())
	assert.Nil(t, jwtToken)
}
