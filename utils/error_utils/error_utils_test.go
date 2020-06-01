package error_utils

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewApiErrorSuccessful(t *testing.T) {

	//arrange
	internalStatusCode := 2000
	message := "Already exists"

	//act
	apiError := NewInternalCustomError(internalStatusCode, message)

	//assert
	assert.NotNil(t, apiError)
	assert.EqualValues(t, message, apiError.ErrorMessage)
	assert.EqualValues(t, http.StatusInternalServerError, apiError.HttpStatusCode)
	assert.EqualValues(t, internalStatusCode, apiError.InternalStatusCode)
}

func TestNewInternalServerErrorSuccessful(t *testing.T) {

	//arrange
	message := "Internal Server Error"

	//act
	apiError := NewInternalServerError(message)

	//assert
	assert.NotNil(t, apiError)
	assert.EqualValues(t, message, apiError.ErrorMessage)
	assert.EqualValues(t, http.StatusInternalServerError, apiError.HttpStatusCode)
	assert.EqualValues(t, http.StatusInternalServerError, apiError.InternalStatusCode)
}

func TestNewBadRequestErrorSuccessful(t *testing.T) {

	//arrange
	message := "Bad request"

	//act
	apiError := NewBadRequestError(message)

	//assert
	assert.NotNil(t, apiError)
	assert.EqualValues(t, message, apiError.ErrorMessage)
	assert.EqualValues(t, http.StatusBadRequest, apiError.HttpStatusCode)
	assert.EqualValues(t, http.StatusBadRequest, apiError.InternalStatusCode)
}

func TestUnauthorizedErrorSuccessful(t *testing.T) {

	//arrange
	message := "Unauthorized Error"

	//act
	apiError := NewUnauthorizedError(message)

	//assert
	assert.NotNil(t, apiError)
	assert.EqualValues(t, message, apiError.ErrorMessage)
	assert.EqualValues(t, http.StatusUnauthorized, apiError.HttpStatusCode)
	assert.EqualValues(t, http.StatusUnauthorized, apiError.InternalStatusCode)
}

func TestForbiddenErrorSuccessful(t *testing.T) {

	//arrange
	message := "Forbidden Error"

	//act
	apiError := NewForbiddenError(message)

	//assert
	assert.NotNil(t, apiError)
	assert.EqualValues(t, message, apiError.ErrorMessage)
	assert.EqualValues(t, http.StatusForbidden, apiError.HttpStatusCode)
	assert.EqualValues(t, http.StatusForbidden, apiError.InternalStatusCode)
}
