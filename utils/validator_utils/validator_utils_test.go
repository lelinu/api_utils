package validator_utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIsNotEmptyWithSpacesSuccessful(t *testing.T){
	// arrange
	propName := "PropName"
	input := "  "
	expectedErr := "PropName - Value must not be empty"

	// act
	validator := NewValidator()
	valid := validator.IsNotEmpty(propName, input)

	// assert
	assert.EqualValues(t, false, valid)
	assert.NotNil(t, validator.Err)
	assert.EqualValues(t, expectedErr, validator.Err.Error())
	assert.EqualValues(t, false, validator.IsValid())
}

func TestIsNotEmptyWithBlankInputSuccessful(t *testing.T){
	// arrange
	propName := "PropName"
	input := ""
	expectedErr := "PropName - Value must not be empty"

	// act
	validator := NewValidator()
	valid := validator.IsNotEmpty(propName, input)

	// assert
	assert.EqualValues(t, false, valid)
	assert.NotNil(t, validator.Err)
	assert.EqualValues(t, expectedErr, validator.Err.Error())
	assert.EqualValues(t, false, validator.IsValid())
}

func TestIsNotEmptyValidSuccessful(t *testing.T){
	// arrange
	propName := "PropName"
	input := "Hello world"

	// act
	validator := NewValidator()
	valid := validator.IsNotEmpty(propName, input)

	// assert
	assert.EqualValues(t, true, valid)
	assert.Nil(t, validator.Err)
}

func TestIsAlphaDashWithSpaceSuccessful(t *testing.T){
	// arrange
	propName := "PropName"
	input := "Hell0- World"
	expectedErr := "PropName - Value must be alpha. No spaces allowed, only dashes"

	// act
	validator := NewValidator()
	valid := validator.IsAlphaDash(propName, input)

	// assert
	assert.EqualValues(t, false, valid)
	assert.NotNil(t, validator.Err)
	assert.EqualValues(t, expectedErr, validator.Err.Error())
	assert.EqualValues(t, false, validator.IsValid())
}

func TestIsAlphaDashValidSuccessful(t *testing.T){
	// arrange
	propName := "PropName"
	input := "Hell0-World"

	// act
	validator := NewValidator()
	valid := validator.IsAlphaDash(propName, input)

	// assert
	assert.EqualValues(t, true, valid)
	assert.Nil(t, validator.Err)
}

func TestIsAlphaDashSpaceWithNonAlphaSuccessful(t *testing.T){
	// arrange
	propName := "PropName"
	input := "Hell0- World &7"
	expectedErr := "PropName - Value must be alpha. Only spaces and dashes are allowed"

	// act
	validator := NewValidator()
	valid := validator.IsAlphaDashSpace(propName, input)

	// assert
	assert.EqualValues(t, false, valid)
	assert.NotNil(t, validator.Err)
	assert.EqualValues(t, expectedErr, validator.Err.Error())
	assert.EqualValues(t, false, validator.IsValid())
}

func TestIsAlphaDashSpaceValidSuccessful(t *testing.T){
	// arrange
	propName := "PropName"
	input := "Hell0- World"

	// act
	validator := NewValidator()
	valid := validator.IsAlphaDashSpace(propName, input)

	// assert
	assert.EqualValues(t, true, valid)
	assert.Nil(t, validator.Err)
}

func TestIsValidPasswordLowerCaseMissingSuccessful(t *testing.T){
	// arrange
	propName := "Password"
	input := "H345454%"
	minLen := 5
	maxLen := 10
	expectedErr := "Password - lowercase letter missing"

	// act
	validator := NewValidator()
	valid := validator.IsValidPassword(propName, input, minLen, maxLen)

	// assert
	assert.EqualValues(t, false, valid)
	assert.NotNil(t, validator.Err)
	assert.EqualValues(t, expectedErr, validator.Err.Error())
	assert.EqualValues(t, false, validator.IsValid())
}

func TestIsValidPasswordUpperCaseMissingSuccessful(t *testing.T){
	// arrange
	propName := "Password"
	input := "hell0wrld%"
	minLen := 5
	maxLen := 10
	expectedErr := "Password - uppercase letter missing"

	// act
	validator := NewValidator()
	valid := validator.IsValidPassword(propName, input, minLen, maxLen)

	// assert
	assert.EqualValues(t, false, valid)
	assert.NotNil(t, validator.Err)
	assert.EqualValues(t, expectedErr, validator.Err.Error())
	assert.EqualValues(t, false, validator.IsValid())
}

func TestIsValidPasswordNumberMissingSuccessful(t *testing.T){
	// arrange
	propName := "Password"
	input := "hellowrOld%"
	minLen := 5
	maxLen := 12
	expectedErr := "Password - at least one numeric character required"

	// act
	validator := NewValidator()
	valid := validator.IsValidPassword(propName, input, minLen, maxLen)

	// assert
	assert.EqualValues(t, false, valid)
	assert.NotNil(t, validator.Err)
	assert.EqualValues(t, expectedErr, validator.Err.Error())
	assert.EqualValues(t, false, validator.IsValid())
}

func TestIsValidPasswordSpecialCharMissingSuccessful(t *testing.T){
	// arrange
	propName := "Password"
	input := "Hell0WoRld"
	minLen := 5
	maxLen := 10
	expectedErr := "Password - special character missing"

	// act
	validator := NewValidator()
	valid := validator.IsValidPassword(propName, input, minLen, maxLen)

	// assert
	assert.EqualValues(t, false, valid)
	assert.NotNil(t, validator.Err)
	assert.EqualValues(t, expectedErr, validator.Err.Error())
	assert.EqualValues(t, false, validator.IsValid())
}

func TestIsValidPasswordNotBetweenLengthSuccessful(t *testing.T){
	// arrange
	propName := "Password"
	input := "Hell0WoRld&2020"
	minLen := 5
	maxLen := 10
	expectedErr := "Password - length must be between 5 to 10 characters long"

	// act
	validator := NewValidator()
	valid := validator.IsValidPassword(propName, input, minLen, maxLen)

	// assert
	assert.EqualValues(t, false, valid)
	assert.NotNil(t, validator.Err)
	assert.EqualValues(t, expectedErr, validator.Err.Error())
	assert.EqualValues(t, false, validator.IsValid())
}

func TestIsValidPasswordWithZeroMinMaxLengthsSuccessful(t *testing.T){
	// arrange
	propName := "Password"
	input := "Hell0WoRld&2020"
	minLen := 0
	maxLen := 0
	expectedErr := "Password - max length cannot be zero"

	// act
	validator := NewValidator()
	valid := validator.IsValidPassword(propName, input, minLen, maxLen)

	// assert
	assert.EqualValues(t, false, valid)
	assert.NotNil(t, validator.Err)
	assert.EqualValues(t, expectedErr, validator.Err.Error())
	assert.EqualValues(t, false, validator.IsValid())
}

func TestMaxLengthInvalidSuccessful(t *testing.T){
	// arrange
	propName := "MaxLength"
	input := "Hell0G0LanG&2020"
	maxLen := 10
	expectedErr := "MaxLength - max length is 10"

	// act
	validator := NewValidator()
	valid := validator.MaxLength(propName, input, maxLen)

	// assert
	assert.EqualValues(t, false, valid)
	assert.NotNil(t, validator.Err)
	assert.EqualValues(t, expectedErr, validator.Err.Error())
	assert.EqualValues(t, false, validator.IsValid())
}

func TestMaxLengthValidSuccessful(t *testing.T){
	// arrange
	propName := "MaxLength"
	input := "Hell0G0LanG&2020"
	maxLen := 20

	// act
	validator := NewValidator()
	valid := validator.MaxLength(propName, input, maxLen)

	// assert
	assert.EqualValues(t, true, valid)
	assert.Nil(t, validator.Err)
}

func TestMustBeGreaterThanInvalidSuccessful(t *testing.T){
	// arrange
	propName := "MustBeGreaterThan"
	input := 9
	high := 10
	expectedErr := "MustBeGreaterThan - Value must be greater than 10"

	// act
	validator := NewValidator()
	valid := validator.MustBeGreaterThan(propName, high, input)

	// assert
	assert.EqualValues(t, false, valid)
	assert.NotNil(t, validator.Err)
	assert.EqualValues(t, expectedErr, validator.Err.Error())
	assert.EqualValues(t, false, validator.IsValid())
}

func TestMustBeGreaterThanValidSuccessful(t *testing.T){
	// arrange
	propName := "MustBeGreaterThan"
	input := 11
	high := 10

	// act
	validator := NewValidator()
	valid := validator.MustBeGreaterThan(propName, high, input)

	// assert
	assert.EqualValues(t, true, valid)
	assert.Nil(t, validator.Err)
}

func TestContainsOptionalInvalidSuccessful(t *testing.T){
	// arrange
	propName := "ContainsOptional"
	input := "hello"
	allowedList := make([]string, 0)
	allowedList = append(allowedList, "Golang")
	expectedErr := "ContainsOptional - Value is not in the allowed list: Golang"

	// act
	validator := NewValidator()
	valid := validator.Contains(propName, input, allowedList, true)

	// assert
	assert.EqualValues(t, false, valid)
	assert.NotNil(t, validator.Err)
	assert.EqualValues(t, expectedErr, validator.Err.Error())
	assert.EqualValues(t, false, validator.IsValid())
}

func TestContainsOptionalValidSuccessful(t *testing.T){
	// arrange
	propName := "ContainsOptional"
	input := "Golang"
	allowedList := make([]string, 0)
	allowedList = append(allowedList, "Golang")

	// act
	validator := NewValidator()
	valid := validator.Contains(propName, input, allowedList, true)

	// assert
	assert.EqualValues(t, true, valid)
	assert.Nil(t, validator.Err)
}

func TestContainsOptionalEmptyValueSuccessful(t *testing.T){
	// arrange
	propName := "ContainsOptional"
	input := ""
	allowedList := make([]string, 0)
	allowedList = append(allowedList, "Golang")

	// act
	validator := NewValidator()
	valid := validator.Contains(propName, input, allowedList, true)

	// assert
	assert.EqualValues(t, true, valid)
	assert.Nil(t, validator.Err)
}

func TestContainsRequiredEmptyValueSuccessful(t *testing.T){
	// arrange
	propName := "ContainsOptional"
	input := ""
	allowedList := make([]string, 0)
	allowedList = append(allowedList, "Golang")
	expectedErr := "ContainsOptional - Value must not be empty"

	// act
	validator := NewValidator()
	valid := validator.Contains(propName, input, allowedList, false)

	// assert
	assert.EqualValues(t, false, valid)
	assert.NotNil(t, validator.Err)
	assert.EqualValues(t, expectedErr, validator.Err.Error())
	assert.EqualValues(t, false, validator.IsValid())
}

func TestIsFixedLengthInvalidSuccessful(t *testing.T){
	// arrange
	propName := "FixedLength"
	input := "hello"
	fixedLen := 4
	expectedErr := "FixedLength - Value length must be of 4"

	// act
	validator := NewValidator()
	valid := validator.IsFixedLength(propName, input, fixedLen)

	// assert
	assert.EqualValues(t, false, valid)
	assert.NotNil(t, validator.Err)
	assert.EqualValues(t, expectedErr, validator.Err.Error())
	assert.EqualValues(t, false, validator.IsValid())
}

func TestIsFixedLengthValidSuccessful(t *testing.T){
	// arrange
	propName := "FixedLength"
	input := "hello"
	fixedLen := 5

	// act
	validator := NewValidator()
	valid := validator.IsFixedLength(propName, input, fixedLen)

	// assert
	assert.EqualValues(t, true, valid)
	assert.Nil(t, validator.Err)
}

func TestIsMinLengthInvalidSuccessful(t *testing.T){
	// arrange
	propName := "MinLength"
	input := "hi"
	minLength := 3
	expectedErr := "MinLength - Value length must be greater or equal to 3"

	// act
	validator := NewValidator()
	valid := validator.IsMinLength(propName, input, minLength)

	// assert
	assert.EqualValues(t, false, valid)
	assert.NotNil(t, validator.Err)
	assert.EqualValues(t, expectedErr, validator.Err.Error())
	assert.EqualValues(t, false, validator.IsValid())
}

func TestIsMinLengthValidSuccessful(t *testing.T){
	// arrange
	propName := "MinLength"
	input := "hello"
	minLength := 5

	// act
	validator := NewValidator()
	valid := validator.IsMinLength(propName, input, minLength)

	// assert
	assert.EqualValues(t, true, valid)
	assert.Nil(t, validator.Err)
}

func TestIsNotNilInvalidSuccessful(t *testing.T){
	// arrange
	propName := "NotNil"
	var input interface{}
	expectedErr := "NotNil - Value must not be nil"

	// act
	validator := NewValidator()
	valid := validator.IsNotNil(propName, input)

	// assert
	assert.EqualValues(t, false, valid)
	assert.NotNil(t, validator.Err)
	assert.EqualValues(t, expectedErr, validator.Err.Error())
	assert.EqualValues(t, false, validator.IsValid())
}

func TestIsNotNilValidSuccessful(t *testing.T){
	// arrange
	propName := "NotNil"
	input := map[string]interface{}{"value": "value"}

	// act
	validator := NewValidator()
	valid := validator.IsNotNil(propName, input)

	// assert
	assert.EqualValues(t, true, valid)
	assert.Nil(t, validator.Err)
}

func TestIsValidEmailInvalidSuccessful(t *testing.T){
	// arrange
	propName := "ValidEmail"
	var input = "hello.gmail.com"
	expectedErr := "ValidEmail - Value is not a valid email address"

	// act
	validator := NewValidator()
	valid := validator.IsValidEmail(propName, input)

	// assert
	assert.EqualValues(t, false, valid)
	assert.NotNil(t, validator.Err)
	assert.EqualValues(t, expectedErr, validator.Err.Error())
	assert.EqualValues(t, false, validator.IsValid())
}

func TestIsValidEmailValidSuccessful(t *testing.T){
	// arrange
	propName := "ValidEmail"
	var input = "Hello1234_@gmail.com"

	// act
	validator := NewValidator()
	valid := validator.IsValidEmail(propName, input)

	// assert
	assert.EqualValues(t, true, valid)
	assert.Nil(t, validator.Err)
}

func TestIsValidURLWithoutHttpsInvalidSuccessful(t *testing.T){
	// arrange
	propName := "ValidURL"
	input := "hi.gmail.com"
	expectedErr := "ValidURL - Value is not a valid url"

	// act
	validator := NewValidator()
	valid := validator.IsValidURL(propName, input, false)

	// assert
	assert.EqualValues(t, false, valid)
	assert.NotNil(t, validator.Err)
	assert.EqualValues(t, expectedErr, validator.Err.Error())
	assert.EqualValues(t, false, validator.IsValid())
}

func TestIsValidURLWithoutHttpsValidSuccessful(t *testing.T){
	// arrange
	propName := "ValidURL"
	input := "http://hi.gmail.com"

	// act
	validator := NewValidator()
	valid := validator.IsValidURL(propName, input, false)

	// assert
	assert.EqualValues(t, true, valid)
	assert.Nil(t, validator.Err)
}

func TestIsValidURLWithHttpsInvalidSuccessful(t *testing.T){
	// arrange
	propName := "ValidURL"
	input := "http://hi.gmail.com"
	expectedErr := "ValidURL - Value is not a valid https url"

	// act
	validator := NewValidator()
	valid := validator.IsValidURL(propName, input, true)

	// assert
	assert.EqualValues(t, false, valid)
	assert.NotNil(t, validator.Err)
	assert.EqualValues(t, expectedErr, validator.Err.Error())
	assert.EqualValues(t, false, validator.IsValid())
}

func TestIsValidURLWithHttpsValidSuccessful(t *testing.T){
	// arrange
	propName := "ValidURL"
	input := "https://hi.gmail.com"

	// act
	validator := NewValidator()
	valid := validator.IsValidURL(propName, input, true)

	// assert
	assert.EqualValues(t, true, valid)
	assert.Nil(t, validator.Err)
}

func TestDateMustBeBeforeInvalidSuccessful(t *testing.T){
	// arrange
	propName := "DateMustBeBefore"
	dateFormat :=  time.RFC3339
	input, _ := time.Parse(dateFormat, "2020-05-04T12:42:51Z")
	high, _ :=  time.Parse(dateFormat, "2020-05-04T12:31:51Z")
	expectedErr := "DateMustBeBefore - Value must be before than 2020-05-04 12:31:51 +0000 UTC"

	// act
	validator := NewValidator()
	valid := validator.DateMustBeBefore(propName, input, high)

	// assert
	assert.EqualValues(t, false, valid)
	assert.NotNil(t, validator.Err)
	assert.EqualValues(t, expectedErr, validator.Err.Error())
	assert.EqualValues(t, false, validator.IsValid())
}

func TestDateMustBeBeforeValidSuccessful(t *testing.T){
	// arrange
	propName := "DateMustBeBefore"
	dateFormat :=  time.RFC3339
	input, _ := time.Parse(dateFormat, "2020-05-04T12:30:51Z")
	high, _ :=  time.Parse(dateFormat, "2020-05-04T12:31:51Z")

	// act
	validator := NewValidator()
	valid := validator.DateMustBeBefore(propName, input, high)

	// assert
	assert.EqualValues(t, true, valid)
	assert.Nil(t, validator.Err)
}

func TestDateMustNotBeInFutureInvalidSuccessful(t *testing.T){
	// arrange
	propName := "DateMustBeBefore"
	input := time.Now().UTC().Add(time.Second * 2)
	expectedErr := "DateMustBeBefore - Value cannot be in the future"

	// act
	validator := NewValidator()
	valid := validator.DateMustNotBeInFuture(propName, input)

	// assert
	assert.EqualValues(t, false, valid)
	assert.NotNil(t, validator.Err)
	assert.EqualValues(t, expectedErr, validator.Err.Error())
	assert.EqualValues(t, false, validator.IsValid())
}

func TestDateMustNotBeInFutureValidSuccessful(t *testing.T){
	// arrange
	propName := "DateMustBeBefore"
	dateFormat :=  time.RFC3339
	input, _ := time.Parse(dateFormat, "2020-03-04T12:30:51Z")

	// act
	validator := NewValidator()
	valid := validator.DateMustNotBeInFuture(propName, input)

	// assert
	assert.EqualValues(t, true, valid)
	assert.Nil(t, validator.Err)
}

