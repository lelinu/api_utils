package jwe

import (
	"errors"
	"fmt"
	"github.com/lelinu/api_utils/utils/error_utils"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
	"strings"
	"time"
)

//Service struct
type Service struct {
	timeFunc            func() time.Time
	timeout             time.Duration
	maxRefresh          time.Duration
	encryptionAlgorithm string
	encryptionKey       []byte
	issuer 				string
}

// NewService this method will return a new instance of JweService
// Currently supports JWE encryption only
func NewService(encryptionAlgorithm string, encryptionKey string, issuer string,
	timeoutInHours time.Duration, maxRefreshInHours time.Duration) (*Service, *error_utils.ApiError) {

	var service = &Service{}
	err := service.init(encryptionAlgorithm, encryptionKey, issuer, timeoutInHours, maxRefreshInHours)
	if err != nil {
		return nil, err
	}

	return service, nil
}

//init will initialize defaults
func (a *Service) init(encryptionAlgorithm string, encryptionKey string, issuer string,
	timeoutInHours time.Duration, maxRefreshInHours time.Duration) *error_utils.ApiError {

	// validations
	// validate encryption algorithm
	if err := a.validateEncryptionAlgorithm(encryptionAlgorithm); err != nil {
		return error_utils.NewBadRequestError(fmt.Sprintf("Auth: %v", err.Error()))
	}

	// validate issuer
	if strings.TrimSpace(issuer) == ""{
		return error_utils.NewBadRequestError("Auth: issuer cannot be empty")
	}

	// validate max refresh
	if maxRefreshInHours <= 0{
		return error_utils.NewBadRequestError("Auth: max refresh should be greater than 0")
	}

	// validate jwt secret key
	if len(strings.TrimSpace(encryptionKey)) == 0 {
		return error_utils.NewBadRequestError("Auth: encryption key cannot be empty")
	}

	// validations

	a.encryptionAlgorithm = encryptionAlgorithm
	a.encryptionKey = []byte(encryptionKey)
	a.issuer = issuer
	a.maxRefresh = time.Hour * maxRefreshInHours

	// set defaults
	a.timeFunc = time.Now
	if a.timeout <= 0 {
		a.timeout = time.Hour * timeoutInHours
	}
	// set defaults

	return nil
}

//GenerateJweToken will generate a new jwe token
func (a *Service) GenerateJweToken(customClaims map[string]interface{}) (string, *time.Time, *error_utils.ApiError) {

	enc, err := jose.NewEncrypter(
		jose.ContentEncryption(a.encryptionAlgorithm),
		jose.Recipient{Algorithm: jose.DIRECT, Key: a.encryptionKey},
		(&jose.EncrypterOptions{}).WithType("JWT"),
	)
	if err != nil {
		return "", nil, error_utils.NewInternalServerError(err.Error())
	}

	expire := a.timeFunc().UTC().Add(a.timeout)

	claims := map[string]interface{} { }
	claims["exp"] = expire.Unix()
	claims["orig_iat"] = a.timeFunc().Unix()
	claims["iss"] = a.issuer

	if customClaims != nil {
		for key, value := range customClaims {
			claims[key] = value
		}
	}

	token, err := jwt.Encrypted(enc).Claims(claims).CompactSerialize()
	if err != nil {
		return "", nil, error_utils.NewInternalServerError(err.Error())
	}

	return token, &expire, nil
}

//RefreshToken will generate a token based on original token
func (a *Service) RefreshJweToken(token string) (string, *time.Time, *error_utils.ApiError){

	claims, apiErr := a.ValidateJweToken(token)
	if apiErr != nil {
		return "", nil, apiErr
	}

	enc, err := jose.NewEncrypter(
		jose.ContentEncryption(a.encryptionAlgorithm),
		jose.Recipient{Algorithm: jose.DIRECT, Key: a.encryptionKey},
		(&jose.EncrypterOptions{}).WithType("JWT"),
	)
	if err != nil {
		return "", nil, error_utils.NewInternalServerError(err.Error())
	}

	newClaims :=  map[string]interface{} {}
	for key := range claims {
		newClaims[key] = claims[key]
	}

	expire := a.timeFunc().UTC().Add(a.timeout)
	newClaims["exp"] = expire.Unix()
	newClaims["orig_iat"] = a.timeFunc().Unix()

	token, err = jwt.Encrypted(enc).Claims(newClaims).CompactSerialize()
	if err != nil {
		return "", nil, error_utils.NewInternalServerError(err.Error())
	}

	return token, &expire, nil
}

//ValidateJweToken will validate the Jwe token
func (a *Service) ValidateJweToken(token string) (map[string]interface{}, *error_utils.ApiError) {

	// parse token string
	claims, err := a.parseTokenString(token)
	if err != nil {
		return nil, error_utils.NewUnauthorizedError(err.Error())
	}

	fmt.Printf("claims are %v", claims)

	// validate dates
	if claims["orig_iat"] == nil {
		return nil, error_utils.NewUnauthorizedError("Orig Iat is missing")
	}

	// try convert to float64
	if _, ok := claims["orig_iat"].(float64); !ok {
		return nil, error_utils.NewUnauthorizedError("Orig Iat must be float64 format")
	}

	// get value and validate
	origIat := int64(claims["orig_iat"].(float64))
	if origIat < a.timeFunc().Add(-a.maxRefresh).Unix() {
		fmt.Println("dhalt awn 1...")
		return nil, error_utils.NewUnauthorizedError("Token is expired")
	}

	// check if exp exists in map
	if claims["exp"] == nil {
		return nil, error_utils.NewUnauthorizedError("Exp is missing")
	}

	// try convert to float 64
	if _, ok := claims["exp"].(float64); !ok {
		return nil, error_utils.NewUnauthorizedError("Exp must be float64 format")
	}

	// get value and validate
	exp := int64(claims["exp"].(float64))
	if exp < a.timeFunc().Unix(){
		fmt.Println("dhalt awn 2...")
		return nil, error_utils.NewUnauthorizedError("Token is expired")
	}
	// validate dates

	// validate issuer
	// check if iss exists in map
	if claims["iss"] == nil {
		return nil, error_utils.NewUnauthorizedError("Iss is missing")
	}

	// try convert to string
	if _, ok := claims["iss"].(string); !ok {
		return nil, error_utils.NewUnauthorizedError("Iss must be string format")
	}

	// get value and validate
	issuer := claims["iss"]
	if issuer != a.issuer{
		return nil, error_utils.NewUnauthorizedError("Invalid issuer")
	}
	// validate issuer

	return claims, nil
}

//parseTokenString will parse the token claims
func (a *Service) parseTokenString(token string) (map[string]interface{}, error) {
	tok, err := jwt.ParseEncrypted(token)
	if err != nil {
		return nil, err
	}

	claims := map[string]interface{} {}
	if err := tok.Claims(a.encryptionKey, &claims); err != nil {
		return nil, err
	}
	return claims, nil
}

//validateEncryptionAlgorithm will validate the encryption algorithm
func (a *Service) validateEncryptionAlgorithm(encryptionAlgorithm string) error {

	// convert to Jose content type
	encAlgorithm := jose.ContentEncryption(encryptionAlgorithm)

	if encAlgorithm == jose.A128CBC_HS256 ||
		encAlgorithm == jose.A192CBC_HS384 ||
		encAlgorithm == jose.A256CBC_HS512 ||
		encAlgorithm == jose.A128GCM ||
		encAlgorithm == jose.A192GCM ||
		encAlgorithm == jose.A256GCM {
		return nil
	}

	return errors.New("invalid encryption algorithm")
}
