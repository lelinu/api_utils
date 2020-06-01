package parameterstore

import (
	"errors"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	funcGetParamsByPath func(input *ssm.GetParametersByPathInput) (*ssm.GetParametersByPathOutput, error)
	keyPath             = "keypath"
	region              = "eu-west-1"
	withDecryption      = false
)

type SsmServiceMock struct{}

func (a SsmServiceMock) GetParametersByPath(input *ssm.GetParametersByPathInput) (*ssm.GetParametersByPathOutput, error) {
	return funcGetParamsByPath(input)
}

func TestInitInvalidKeyPath(t *testing.T){

	service, apiErr := NewService("", region, withDecryption)

	assert.NotNil(t, apiErr)
	assert.EqualValues(t, "Parameter Store: Key path is empty", apiErr.ErrorMessage)
	assert.Nil(t, service)
}

func TestInitInvalidRegion(t *testing.T){

	service, apiErr := NewService(keyPath, "", withDecryption)

	assert.NotNil(t, apiErr)
	assert.EqualValues(t, "Parameter Store: Region is empty", apiErr.ErrorMessage)
	assert.Nil(t, service)
}

//TestGetParametersByPathValid
func TestGetParametersByPathValid(t *testing.T) {

	serviceMock := SsmServiceMock{}
	funcGetParamsByPath = func(input *ssm.GetParametersByPathInput) (*ssm.GetParametersByPathOutput, error) {
		ssmResult := ssm.Parameter{
			Name:  newString("name"),
			Value: newString("value"),
		}

		var params []*ssm.Parameter
		params = append(params, &ssmResult)

		return &ssm.GetParametersByPathOutput{
			Parameters: params,
		}, nil
	}

	service, apiErr := NewServiceMock(keyPath, region, withDecryption, serviceMock)
	assert.Nil(t, apiErr)
	assert.NotNil(t, service)

	res, err := service.GetParametersByPath()

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.NotNil(t, res["name"])
	assert.EqualValues(t, "value", res["name"])
}

//TestGetParametersByPathInValid
func TestGetParametersByPathInValid(t *testing.T) {

	serviceMock := SsmServiceMock{}
	funcGetParamsByPath = func(input *ssm.GetParametersByPathInput) (*ssm.GetParametersByPathOutput, error) {
		return nil, errors.New("path not found")
	}

	service, apiErr := NewServiceMock(keyPath, region, withDecryption, serviceMock)
	assert.Nil(t, apiErr)
	assert.NotNil(t, service)

	res, err := service.GetParametersByPath()

	assert.NotNil(t, err)
	assert.EqualValues(t, "parameterStore: GetParametersByPath : An error had occurred in parameters by path - path not found", err.ErrorMessage)
	assert.Nil(t, res)
}

func newString(value string) *string {
	return &value
}
