package parameterstore

import (
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/lelinu/api_utils/utils/error_utils"
)

//IService interface
type IService interface {
	GetParametersByPath() (map[string]interface{}, *error_utils.ApiError)
}

//ISsmService this is used generally for mocking
type ISsmService interface {
	GetParametersByPath(input *ssm.GetParametersByPathInput) (*ssm.GetParametersByPathOutput, error)
}