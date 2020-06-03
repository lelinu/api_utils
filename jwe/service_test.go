package jwe

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var
(
	encryptionAlgorithm string
	encryptionKey       string
	issuer              string
	customClaims        map[string]interface{}
)

func init() {
	encryptionAlgorithm = "A256GCM"
	encryptionKey = "s4IIq9lQm2SKBlJoHAWzkRGSNaPCLZw2"
	issuer = "lelinu"
	customClaims = map[string]interface{}{"id": 1, "role": "user", "merchantID": "1234546"}
}

func TestToken(t *testing.T){
	service, _ := NewService("A256GCM", "rbCdjdybxncSb6drqd9bvBcfembQWbkv", "https://golelinu.com", 1, 1)
	var token = "eyJhbGciOiJkaXIiLCJlbmMiOiJBMjU2R0NNIiwidHlwIjoiSldUIn0..retSH04y4B7BcseU.U2-yKi3bntdOPidgof_31BREs7UkfxW-FiDKl7t6d-2mNjb9K9tGkOrWckFrnsoxG7Xe2g8uxq_WNfxbhmTCFswpJMCOamUe_9YQqceHNEkp_y_3qjU052_p_Xlmnjjlgf9YQ72EQE20e_46vatDDkN9XgtvykRL1_X2y97ZQ6Wer_bQklHXNnEo5NxqNvJEWXF46ymUI9CZlDR__iZBPH5csIvhXnD0jxVCeTwoZ1g4B9lqgVocnO6blpCzxpQNSJm2TcfAFEENRgVAANVM7oST94Nhz2BR-Q5rEISfRp0RXJnFES4-bXs2Un1aPG39QLPRC5MeEDtKZ3Ru9PaQEOBHX8RQqe5Ym6lMqUbU3kL7MG6XOaess5JlvYuHop9oUS1QAo12VwbsTsqRydYyjvKAW3XQNEi98YFQn7PVxm-caSp1BXinNDwj64o1WjEV1RSR1sgYHKZtMuF5OAS4Hmc9oE5qe1xtw1ZJRwTtKNWskxljmcPL4K9p7xvzHARFJU7koqRqJeP3IIX3sYNtQj3ri0L_OzlLSUTKoDwo4W5IQpb34t_H1KbNxhFWh0pxFB6iNsMYNqge8gjNwzr07HRLjkBi2hSpJ1_pVCmb9mXZPPIRmP50-iiAG5VKrLtik7hw4wBsuPhlnZ7qF6cR3X6wSrwMPtdxApuU45R37_j8_E-BXyE9cqv6AofO-j98flo4IppNBjNRkkFjM5_WP330f5f1SMw.nUhln8D77nGXtdaaedO2DA"
	_, err := service.ValidateJweToken(token)
	fmt.Printf("Error is %v", err)
}

func TestNewServiceInvalidAlgorithm(t *testing.T) {

	_, err := NewService("invalid-algo", encryptionKey, issuer, 1, 1)
	assert.NotNil(t, err)
	assert.EqualValues(t, "Auth: invalid encryption algorithm", err.ErrorMessage)
}

//TestInitWithInvalidAlgorithm
func TestInitWithInvalidAlgorithm(t *testing.T) {

	service := Service{}
	err := service.init("invalid-algo", encryptionKey, issuer, 1, 1)

	assert.NotNil(t, err)
	assert.EqualValues(t, "Auth: invalid encryption algorithm", err.ErrorMessage)
}

//TestInitWithInvalidIssuer
func TestInitWithInvalidIssuer(t *testing.T) {

	service := Service{}
	err := service.init(encryptionAlgorithm, encryptionKey, "", 1, 1)

	assert.NotNil(t, err)
	assert.EqualValues(t, "Auth: issuer cannot be empty", err.ErrorMessage)
}

//TestInitWithInvalidMaxRefresh
func TestInitWithInvalidMaxRefresh(t *testing.T) {

	service := Service{}
	err := service.init(encryptionAlgorithm, encryptionKey, issuer, 1, 0)

	assert.NotNil(t, err)
	assert.EqualValues(t, "Auth: max refresh should be greater than 0", err.ErrorMessage)
}

//TestGenerateJweTokenWithInvalidEncryptionKeyLength
func TestGenerateJweTokenWithInvalidEncryptionKeyLength(t *testing.T) {
	// arrange
	service, err := NewService(encryptionAlgorithm, "invalid-enc", issuer, 1, 1)
	assert.Nil(t, err)
	assert.NotNil(t, service)

	// act
	token, exp, err := service.GenerateJweToken(nil)

	// assert
	assert.NotNil(t, err)
	assert.EqualValues(t, "", token)
	assert.Nil(t, exp)
}

//TestGenerateJweTokenValidWithoutCustomClaims
func TestGenerateJweTokenValidWithoutCustomClaims(t *testing.T) {

	// arrange
	service, err := NewService(encryptionAlgorithm, encryptionKey, issuer, 1, 1)
	assert.Nil(t, err)
	assert.NotNil(t, service)

	// act
	token, exp, err := service.GenerateJweToken(nil)

	// assert
	assert.NotNil(t, token)
	assert.NotNil(t, exp)
	assert.Greater(t, exp.Unix(), time.Now().UTC().Unix())
	assert.Nil(t, err)
}

//TestGenerateJweTokenValidWithCustomClaims
func TestGenerateJweTokenValidWithCustomClaims(t *testing.T) {

	// arrange
	service, err := NewService(encryptionAlgorithm, encryptionKey, issuer, 1, 1)
	assert.Nil(t, err)
	assert.NotNil(t, service)

	// act
	token, exp, err := service.GenerateJweToken(customClaims)

	// assert
	assert.NotNil(t, token)
	assert.NotNil(t, exp)
	assert.Greater(t, exp.Unix(), time.Now().UTC().Unix())
	assert.Nil(t, err)
}

//TestValidateJweTokenInvalidIssuer
func TestValidateJweTokenInvalidIssuer(t *testing.T) {
	// generate with lelinu as issuer
	gService, err := NewService(encryptionAlgorithm, encryptionKey, issuer, 1, 1)
	assert.Nil(t, err)
	assert.NotNil(t, gService)

	token, _, err := gService.GenerateJweToken(customClaims)
	assert.Nil(t, err)
	assert.NotEqual(t, "", token)

	// validate with test as issuer
	vService, err := NewService(encryptionAlgorithm, encryptionKey, "test", 1, 1)
	assert.Nil(t, err)
	assert.NotNil(t, vService)

	claims, err := vService.ValidateJweToken(token)
	assert.NotNil(t, err)
	assert.EqualValues(t, "Invalid issuer", err.ErrorMessage)
	assert.Nil(t, claims)
}

//TestValidateJweTokenExpiredWithoutCustomClaims
func TestValidateJweTokenExpiredWithoutCustomClaims(t *testing.T) {

	// arrange
	service, err := NewService(encryptionAlgorithm, encryptionKey, issuer, 1, 1)
	token := "eyJhbGciOiJkaXIiLCJlbmMiOiJBMjU2R0NNIiwidHlwIjoiSldUIn0..uOSuqquOI4DWfg74.EDQVhTvy67oEkO3NKjpBmX7dTxa-vI2CjVeVEv3ueG7E3fSm7imYk1hemEbJPjA5B9Cm9BOqlg.aKwfVmc1LcCUlQ_lMrE_VA"

	assert.Nil(t, err)
	assert.NotNil(t, service)

	// act
	claims, err := service.ValidateJweToken(token)

	// assert
	assert.NotNil(t, err)
	assert.EqualValues(t, "Token is expired", err.ErrorMessage)
	assert.Nil(t, claims)
}

//TestValidateJweTokenExpiredWithCustomClaims
func TestValidateJweTokenExpiredWithCustomClaims(t *testing.T) {

	// arrange
	service, err := NewService(encryptionAlgorithm, encryptionKey, issuer, 1, 1)
	token := "eyJhbGciOiJkaXIiLCJlbmMiOiJBMjU2R0NNIiwidHlwIjoiSldUIn0..xPUOeT3A_69oIMAr.sL9jO2nh-EaKQ2J_5QPe6UqK8FbyXtJPNlvjRINcz_pZW7dJid7tZ1mPUz5fznYVLolUgKm7ogcu_oJNd6AKO6lJObo9oFgh-MHToobg2PvbohPWZLGI3bo4hKrfZY3hQfFo.t_Ixd1yPn-g47InC1srJmA"

	assert.Nil(t, err)
	assert.NotNil(t, service)

	// act
	claims, err := service.ValidateJweToken(token)

	// assert
	assert.NotNil(t, err)
	assert.EqualValues(t, "Token is expired", err.ErrorMessage)
	assert.Nil(t, claims)
}

//TestValidateJweTokenValidWithCustomClaims
func TestValidateJweTokenValidWithCustomClaims(t *testing.T) {

	service, err := NewService(encryptionAlgorithm, encryptionKey, issuer, 1, 1)
	assert.Nil(t, err)
	assert.NotNil(t, service)

	// act
	token, exp, err := service.GenerateJweToken(customClaims)
	assert.NotNil(t, token)
	assert.NotNil(t, exp)
	assert.Greater(t, exp.Unix(), time.Now().UTC().Unix())
	assert.Nil(t, err)

	// act
	claims, err := service.ValidateJweToken(token)

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

//TestValidateJweTokenValidWithoutCustomClaims
func TestValidateJweTokenValidWithoutCustomClaims(t *testing.T) {

	service, err := NewService(encryptionAlgorithm, encryptionKey, issuer, 1, 1)
	assert.Nil(t, err)
	assert.NotNil(t, service)

	// act
	token, exp, err := service.GenerateJweToken(nil)
	assert.NotNil(t, token)
	assert.NotNil(t, exp)
	assert.Greater(t, exp.Unix(), time.Now().UTC().Unix())
	assert.Nil(t, err)

	// act
	claims, err := service.ValidateJweToken(token)

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

//TestValidateJweRefreshTokenExpiredWithoutCustomClaims
func TestValidateJweRefreshTokenExpiredWithoutCustomClaims(t *testing.T) {

	// arrange
	service, err := NewService(encryptionAlgorithm, encryptionKey, issuer, 1, 1)
	token := "eyJhbGciOiJkaXIiLCJlbmMiOiJBMjU2R0NNIiwidHlwIjoiSldUIn0..uOSuqquOI4DWfg74.EDQVhTvy67oEkO3NKjpBmX7dTxa-vI2CjVeVEv3ueG7E3fSm7imYk1hemEbJPjA5B9Cm9BOqlg.aKwfVmc1LcCUlQ_lMrE_VA"

	assert.Nil(t, err)
	assert.NotNil(t, service)

	// act
	token, exp, err := service.RefreshJweToken(token)

	// assert
	assert.NotNil(t, err)
	assert.EqualValues(t, "Token is expired", err.ErrorMessage)
	assert.Nil(t, exp)
	assert.EqualValues(t, "", token)
}

//TestGenerateJweRefreshTokenWithInvalidEncryptionKeyLength
func TestGenerateJweRefreshTokenWithInvalidEncryptionKeyLength(t *testing.T) {
	// arrange
	service, err := NewService(encryptionAlgorithm, "s4IIq9lQm2SKBlJoHAWzkRGSNaP", issuer, 1, 1)
	token := "eyJhbGciOiJkaXIiLCJlbmMiOiJBMjU2R0NNIiwidHlwIjoiSldUIn0..uOSuqquOI4DWfg74.EDQVhTvy67oEkO3NKjpBmX7dTxa-vI2CjVeVEv3ueG7E3fSm7imYk1hemEbJPjA5B9Cm9BOqlg.aKwfVmc1LcCUlQ_lMrE_VA"
	assert.Nil(t, err)
	assert.NotNil(t, service)

	// act
	token, exp, err := service.RefreshJweToken(token)

	// assert
	assert.NotNil(t, err)
	assert.EqualValues(t, "", token)
	assert.Nil(t, exp)
}

//TestRefreshTokenValidWithoutCustomClaims
func TestRefreshTokenValidWithoutCustomClaims(t *testing.T) {

	service, err := NewService(encryptionAlgorithm, encryptionKey, issuer, 1, 1)
	assert.Nil(t, err)
	assert.NotNil(t, service)

	orToken, exp, err := service.GenerateJweToken(nil)

	assert.NotNil(t, orToken)
	assert.NotNil(t, exp)
	assert.Greater(t, exp.Unix(), time.Now().UTC().Unix())
	assert.Nil(t, err)

	reToken, reExp, err := service.RefreshJweToken(orToken)
	assert.NotNil(t, reToken)
	assert.NotEqual(t, orToken, reToken)
	assert.NotNil(t, reExp)
	assert.Greater(t, reExp.Unix(), time.Now().UTC().Unix())
	assert.Nil(t, err)
}

//TestRefreshTokenValidWithCustomClaims
func TestRefreshTokenValidWithCustomClaims(t *testing.T) {

	service, err := NewService(encryptionAlgorithm, encryptionKey, issuer, 1, 1)
	assert.Nil(t, err)
	assert.NotNil(t, service)

	orToken, exp, err := service.GenerateJweToken(customClaims)

	assert.NotNil(t, orToken)
	assert.NotNil(t, exp)
	assert.Greater(t, exp.Unix(), time.Now().UTC().Unix())
	assert.Nil(t, err)

	reToken, reExp, err := service.RefreshJweToken(orToken)
	assert.NotNil(t, reToken)
	assert.NotEqual(t, orToken, reToken)
	assert.NotNil(t, reExp)
	assert.Greater(t, reExp.Unix(), time.Now().UTC().Unix())
	assert.Nil(t, err)
}

//TestRefreshTokenValidWithoutClaims
func TestRefreshTokenAndValidateValidWithCustomClaims(t *testing.T) {

	service, err := NewService(encryptionAlgorithm, encryptionKey, issuer, 1, 1)
	assert.Nil(t, err)
	assert.NotNil(t, service)

	orToken, exp, err := service.GenerateJweToken(customClaims)
	assert.NotNil(t, orToken)
	assert.NotNil(t, exp)
	assert.Greater(t, exp.Unix(), time.Now().UTC().Unix())
	assert.Nil(t, err)

	reToken, reExp, err := service.RefreshJweToken(orToken)
	assert.NotNil(t, reToken)
	assert.NotEqual(t, orToken, reToken)
	assert.NotNil(t, reExp)
	assert.Greater(t, reExp.Unix(), time.Now().UTC().Unix())
	assert.Nil(t, err)

	claims, err := service.ValidateJweToken(reToken)

	assert.Nil(t, err)
	assert.NotNil(t, claims)
	assert.NotNil(t, claims["exp"])
	assert.NotNil(t, claims["id"])
	assert.NotNil(t, claims["iss"])
	assert.NotNil(t, claims["merchantID"])
	assert.NotNil(t, claims["orig_iat"])
	assert.NotNil(t, claims["role"])
}

//TestParseTokenStringInvalid
func TestParseTokenStringInvalid(t *testing.T) {
	service, apiError := NewService(encryptionAlgorithm, encryptionKey, issuer, 1, 1)
	token := "hello-world"
	assert.Nil(t, apiError)
	assert.NotNil(t, service)

	claims, err := service.parseTokenString(token)
	assert.NotNil(t, err)
	assert.Nil(t, claims)
	assert.EqualValues(t, "square/go-jose: compact JWE format must have five parts", err.Error())
}

//TestValidateEncryptionAlgorithm
func TestValidateEncryptionAlgorithm(t *testing.T) {
	service, apiError := NewService(encryptionAlgorithm, encryptionKey, issuer, 1, 1)
	algorithm := "hello-world"
	assert.Nil(t, apiError)
	assert.NotNil(t, service)

	err := service.validateEncryptionAlgorithm(algorithm)
	assert.NotNil(t, err)
	assert.EqualValues(t, "invalid encryption algorithm", err.Error())
}
