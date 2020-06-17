package validator_utils

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode"
)

//Validator empty struct
type Validator struct {
	Err error
}

var (
	regexAlphaDash              = regexp.MustCompile(`^[a-zA-Z0-9-]+$`)
	regexAlphaDashSpace         = regexp.MustCompile(`^[a-zA-Z0-9- ]+$`)
	regexEmail                  = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,4}$`)
	regex2SegmentAPIKeyStandard = regexp.MustCompile(`^[a-zA-Z0-9-]+\.[a-zA-Z0-9-]+$`)
)

func NewValidator() *Validator {
	return &Validator{}
}

//IsNotEmpty method to check if input is not empty
func (v *Validator) IsNotEmpty(propertyName string, value string) bool {
	if v.Err != nil {
		return false
	}

	//Trim spaces
	strValue := strings.TrimSpace(value)

	if strValue == "" {
		v.Err = fmt.Errorf("%s - Value must not be empty", propertyName)
		return false
	}
	return true
}

//IsAlphaDash method to check if string is alpha
func (v *Validator) IsAlphaDash(propertyName string, value string) bool {
	if v.Err != nil {
		return false
	}

	// check if input is alpha
	if value != "" && !regexAlphaDash.MatchString(value) {
		v.Err = fmt.Errorf("%s - Value must be alpha. No spaces allowed, only dashes", propertyName)
		return false
	}
	return true
}

//IsAlphaDashSpace method to check if string is alpha
func (v *Validator) IsAlphaDashSpace(propertyName string, value string) bool {
	if v.Err != nil {
		return false
	}

	// check if input is alpha
	if value != "" && !regexAlphaDashSpace.MatchString(value) {
		v.Err = fmt.Errorf("%s - Value must be alpha. Only spaces and dashes are allowed", propertyName)
		return false
	}
	return true
}

//IsValidPassword to validate that password is strong enough
func (v *Validator) IsValidPassword(propertyName string, value string, minLength int, maxLength int) bool {

	if maxLength == 0 {
		v.Err = fmt.Errorf("%s - max length cannot be zero", propertyName)
		return false
	}

	if maxLength < minLength {
		v.Err = fmt.Errorf("%s - max length should be greather than min length", propertyName)
		return false
	}

	var uppercasePresent bool
	var lowercasePresent bool
	var numberPresent bool
	var specialCharPresent bool
	var passLen int
	var errorString string

	for _, ch := range value {
		switch {
		case unicode.IsNumber(ch):
			numberPresent = true
			passLen++
		case unicode.IsUpper(ch):
			uppercasePresent = true
			passLen++
		case unicode.IsLower(ch):
			lowercasePresent = true
			passLen++
		case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
			specialCharPresent = true
			passLen++
		case ch == ' ':
			passLen++
		}
	}
	appendError := func(err string) {
		if len(strings.TrimSpace(errorString)) != 0 {
			errorString += ", " + err
		} else {
			errorString = err
		}
	}
	if !lowercasePresent {
		appendError("lowercase letter missing")
	}
	if !uppercasePresent {
		appendError("uppercase letter missing")
	}
	if !numberPresent {
		appendError("at least one numeric character required")
	}
	if !specialCharPresent {
		appendError("special character missing")
	}
	if !(minLength <= passLen && passLen <= maxLength) {
		appendError(fmt.Sprintf("length must be between %d to %d characters long", minLength, maxLength))
	}

	if len(errorString) != 0 {
		v.Err = fmt.Errorf("%s - "+errorString, propertyName)
		return false
	}
	return true
}

//MaxLength method to check if input is not empty
func (v *Validator) MaxLength(propertyName string, value string, maxLength int) bool {
	if v.Err != nil {
		return false
	}
	if value != "" && len(value) > maxLength {
		v.Err = fmt.Errorf("%s - max length is %d", propertyName, maxLength)
		return false
	}
	return true
}

//MustBeGreaterThan method to check whether value is greater than
func (v *Validator) MustBeGreaterThan(propertyName string, high, value int) bool {
	if v.Err != nil {
		return false
	}
	if value <= high {
		v.Err = fmt.Errorf("%s - Value must be greater than %d", propertyName, high)
		return false
	}
	return true
}

//MustBeGreaterThanFloat64 method to check whether value is greater than
func (v *Validator) MustBeGreaterThanFloat64(propertyName string, high, value float64) bool {
	if v.Err != nil {
		return false
	}
	if value <= high {
		v.Err = fmt.Errorf("%s - Value must be greater than %v", propertyName, high)
		return false
	}
	return true
}

//ContainsList method to check where list is in allowed list
func (v *Validator) ContainsList(propertyName string, list []string, allowedList []string, optional bool) bool {

	if optional == false && len(list) == 0 {
		v.Err = fmt.Errorf("%s - Value must not be empty", propertyName)
		return false
	}

	for _, l := range list {
		val := v.Contains(propertyName, l, allowedList, optional)
		if val == false {
			v.Err = fmt.Errorf("%s - Value is not in the allowed list: %s", propertyName, strings.Join(allowedList, ","))
			return false
		}
	}

	return true
}

//Contains method to check if allowed list contains the inputted value
func (v *Validator) Contains(propertyName string, value string, allowedList []string, optional bool) bool {

	// if optional and value is empty return true
	if optional == true && value == "" {
		return true
	}

	if value == "" {
		v.Err = fmt.Errorf("%s - Value must not be empty", propertyName)
		return false
	}

	for _, n := range allowedList {
		if value == n {
			return true
		}
	}
	v.Err = fmt.Errorf("%s - Value is not in the allowed list: %s", propertyName, strings.Join(allowedList, ","))
	return false
}

//IsFixedLength method to check if input value is in fixed length
func (v *Validator) IsFixedLength(propertyName string, value string, size int) bool {
	if v.Err != nil {
		return false
	}

	if value == "" {
		v.Err = fmt.Errorf("%s - Value must not be empty", propertyName)
		return false
	}

	if len(value) != size {
		v.Err = fmt.Errorf("%s - Value length must be of %d", propertyName, size)
		return false
	}
	return true
}

//IsMinLength method to check if inputted value has minimum length
func (v *Validator) IsMinLength(propertyName string, value string, size int) bool {

	if v.Err != nil {
		return false
	}

	if value == "" {
		v.Err = fmt.Errorf("%s - Value must not be empty", propertyName)
		return false
	}

	if len(value) < size {
		v.Err = fmt.Errorf("%s - Value length must be greater or equal to %d", propertyName, size)
		return false
	}
	return true
}

//IsNotNil method to check if inputted struct is not null
func (v *Validator) IsNotNil(propertyName string, value interface{}) bool {
	if v.Err != nil {
		return false
	}
	if value == nil {
		v.Err = fmt.Errorf("%s - Value must not be nil", propertyName)
		return false
	}
	return true
}

//IsValidEmail method to check if inputted value is an actual email
func (v *Validator) IsValidEmail(propertyName string, email string) bool {

	if v.Err != nil {
		return false
	}

	if email == "" {
		v.Err = fmt.Errorf("%s - Value must not be empty", propertyName)
		return false
	}

	if !regexEmail.MatchString(email) {
		v.Err = fmt.Errorf("%s - Value is not a valid email address", propertyName)
		return false
	}

	return true
}

//IsValidURL method to check if inputted value is a URL. Last parameter enforces a check to be https
func (v *Validator) IsValidURL(propertyName string, inputtedURL string, mustBeHTTPS bool) bool {
	if v.Err != nil {
		return false
	}

	u, err := url.ParseRequestURI(inputtedURL)
	if err != nil {
		v.Err = fmt.Errorf("%s - Value is not a valid url", propertyName)
		return false
	}

	var scheme = u.Scheme

	if mustBeHTTPS {
		if scheme != "https" {
			v.Err = fmt.Errorf("%s - Value is not a valid https url", propertyName)
			return false
		}
	}

	return true
}

//DateMustBeBefore method to check if input is before inputted time
func (v *Validator) DateMustBeBefore(propertyName string, value, high time.Time) bool {
	if v.Err != nil {
		return false
	}
	if value.After(high) {
		v.Err = fmt.Errorf("%s - Value must be before than %v", propertyName, high)
		return false
	}
	return true
}

//DateMustNotBeInFuture method to check whether date is in the future
func (v *Validator) DateMustNotBeInFuture(propertyName string, value time.Time) bool {
	if v.Err != nil {
		return false
	}

	if value.Sub(time.Now().UTC()) > 0 {
		v.Err = fmt.Errorf("%s - Value cannot be in the future", propertyName)
		return false
	}

	return true
}

//IsValidBsonID method to check whether string is a valid mongo bson ID
func (v *Validator) IsValidBsonID(propertyName string, value string) bool {
	if v.Err != nil {
		return false
	}

	if value != "" && !bson.IsObjectIdHex(value) {
		v.Err = fmt.Errorf("%s - Value must be a valid bson object ID", propertyName)
		return false
	}

	return true
}

//IsValid2StepAPIKey method to check whether string is a valid api key containing 2 segments
func (v *Validator) IsValid2SegmentAPIKey(propertyName string, value string) bool {
	if v.Err != nil {
		return false
	}

	// check if input is alpha
	if value != "" && !regex2SegmentAPIKeyStandard.MatchString(value) {
		v.Err = fmt.Errorf("%s - Value must be in the format of abc123.abc123", propertyName)
		return false
	}
	return true
}

//IsNotEmptyStringArray method to check if string array has any elements in it
func (v *Validator) IsNotEmptyStringArray(propertyName string, value []string) bool {
	if v.Err != nil {
		return false
	}

	if len(value) == 0 {
		v.Err = fmt.Errorf("%s - Value must not be an empty array", propertyName)
		return false
	}

	return true
}

//IsValid method to check that all the sub methods called are valid
func (v *Validator) IsValid() bool {
	return v.Err == nil
}

func (v *Validator) Error() error {
	return v.Err
}
