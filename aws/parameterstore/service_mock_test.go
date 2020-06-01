package parameterstore

import (
	"github.com/lelinu/api_utils/utils/error_utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	funcGetParametersByPath func() (map[string]interface{}, *error_utils.ApiError)
)

type parameterStoreMock struct{}

func (a *parameterStoreMock) GetParametersByPath() (map[string]interface{}, *error_utils.ApiError) {
	return funcGetParametersByPath()
}

//TestGetParametersByPathValid
func TestGetParametersByPathMockValid(t *testing.T) {

	service := &parameterStoreMock{}
	funcGetParametersByPath = func() (map[string]interface{}, *error_utils.ApiError) {
		return map[string]interface{} {"value": "hello"}, nil
	}

	res, err := service.GetParametersByPath()

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.NotNil(t, res["value"])
}

//TestGetParametersByPathInValid
func TestGetParametersByPathMockInValid(t *testing.T) {

	service := &parameterStoreMock{}
	funcGetParametersByPath = func() (map[string]interface{}, *error_utils.ApiError) {
		return nil, error_utils.NewInternalServerError("An error")
	}

	res, err := service.GetParametersByPath()

	assert.NotNil(t, err)
	assert.Nil(t, res)
	assert.Nil(t, res["value"])
}
