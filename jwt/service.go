package jwt

import (
	"errors"
	"fmt"
	"github.com/lelinu/api_utils/utils/error_utils"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//Service struct
type Service struct {
	timeFunc         func() time.Time
	timeout          time.Duration
	maxRefresh       time.Duration
	signingAlgorithm string
	jwtSecretKey     []byte
	issuer           string
}

// NewService this method will return a new instance of Jwt Service
// Currently this service only supports HMAC signing
func NewService(signingAlgorithm string, jwtSecretKey string, issuer string,
	timeoutInHours time.Duration, maxRefreshInHours time.Duration) (*Service, *error_utils.ApiError) {
	var service = &Service{}

	err := service.init(signingAlgorithm, jwtSecretKey, issuer, timeoutInHours, maxRefreshInHours)
	if err != nil {
		return nil, err
	}

	return service, nil
}

//init will initialize defaults and validate parameters
func (a *Service) init(signingAlgorithm string, jwtSecretKey string, issuer string,
	timeoutInHours time.Duration, maxRefreshInHours time.Duration) *error_utils.ApiError {

	// validations

	// validate signing algorithm
	if err := a.validateSigningAlgorithm(signingAlgorithm); err != nil{
		return error_utils.NewBadRequestError(fmt.Sprintf("Auth: %v", err.Error()))
	}

	// validate issuer
	if strings.TrimSpace(issuer) == "" {
		return error_utils.NewBadRequestError("Auth: issuer cannot be empty")
	}

	// validate max refresh
	if maxRefreshInHours <= 0 {
		return error_utils.NewBadRequestError("Auth: max refresh should be greater than 0")
	}

	// validate jwt secret key
	if len(strings.TrimSpace(jwtSecretKey)) == 0 {
		return error_utils.NewBadRequestError("Auth: jwt Secret key cannot be empty")
	}
	// validations

	a.signingAlgorithm = signingAlgorithm
	a.jwtSecretKey = []byte(jwtSecretKey)
	a.issuer = issuer
	a.maxRefresh = maxRefreshInHours

	// set defaults
	a.timeFunc = time.Now
	if a.timeout <= 0 {
		a.timeout = time.Hour * timeoutInHours
	}
	// set defaults

	return nil
}

//GenerateJwtToken will generate a new jwt token
func (a *Service) GenerateJwtToken(customClaims map[string]interface{}) (string, *time.Time, *error_utils.ApiError) {

	token := jwt.New(jwt.GetSigningMethod(a.signingAlgorithm))
	claims := token.Claims.(jwt.MapClaims)

	if customClaims != nil {
		for key, value := range customClaims {
			claims[key] = value
		}
	}

	expire := a.timeFunc().UTC().Add(a.timeout)
	claims["exp"] = expire.Unix()
	claims["orig_iat"] = a.timeFunc().Unix()
	claims["iss"] = a.issuer

	tokenString, err := a.signedString(token)
	if err != nil {
		return "", nil, error_utils.NewBadRequestError(err.Error())
	}

	return tokenString, &expire, nil
}

//RefreshJwtToken will refresh a jwt token
func (a *Service) RefreshJwtToken(token string) (string, *time.Time, *error_utils.ApiError) {
	claims, err := a.ValidateJwtToken(token)
	if err != nil {
		return "", nil, err
	}

	// Create the token
	newToken := jwt.New(jwt.GetSigningMethod(a.signingAlgorithm))
	newClaims := newToken.Claims.(jwt.MapClaims)

	for key := range claims {
		newClaims[key] = claims[key]
	}

	expire := a.timeFunc().UTC().Add(a.timeout)
	newClaims["exp"] = expire.Unix()
	newClaims["orig_iat"] = a.timeFunc().Unix()

	tokenString, apiErr := a.signedString(newToken)
	if apiErr != nil {
		return "", nil, err
	}

	return tokenString, &expire, nil
}

//ValidateJwtToken will check if the token is expired
func (a *Service) ValidateJwtToken(token string) (map[string]interface{}, *error_utils.ApiError) {

	parsedToken, err := a.parseTokenString(token)
	if err != nil {
		return nil, error_utils.NewInternalServerError(err.Error())
	}
	claims := parsedToken.Claims.(jwt.MapClaims)

	// validate dates
	origIat := int64(claims["orig_iat"].(float64))
	if origIat < a.timeFunc().Add(-a.maxRefresh).Unix() {
		return nil, error_utils.NewInternalServerError("Token is expired")
	}

	exp := int64(claims["exp"].(float64))
	if exp < a.timeFunc().Unix(){
		return nil, error_utils.NewInternalServerError("Token is expired")
	}
	// validate dates

	// validate issuer
	issuer := claims["iss"]
	if issuer != a.issuer{
		return nil, error_utils.NewInternalServerError("Invalid issuer")
	}
	// validate issuer

	return claims, nil
}

//parseTokenString will parse a token string to a jwt.Token
func (a *Service) parseTokenString(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod(a.signingAlgorithm) != t.Method {
			return nil, errors.New("invalid signature")
		}

		return a.jwtSecretKey, nil
	})
}

//signedString will sign the string
func (a *Service) signedString(token *jwt.Token) (string, error) {
	tokenString, err := token.SignedString(a.jwtSecretKey)
	return tokenString, err
}

//validateSigningAlgorithm will validate the encryption algorithm
func (a *Service) validateSigningAlgorithm(signingAlgorithm string) error {

	if signingAlgorithm == "HS256" ||
		signingAlgorithm == "HS384" ||
		signingAlgorithm == "HS512" {
		return nil
	}

	return errors.New("invalid signing algorithm")
}
