package parameterstore

import (
	"fmt"
	"github.com/lelinu/api_utils/utils/error_utils"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

//Service struct
type Service struct {
	keyPath        string
	region         string
	withDecryption bool
	ssmService 	   ISsmService
}

//NewService this method will return a new instance of ParameterStoreService
func NewService(keyPath string, region string, withDecryption bool) (*Service, *error_utils.ApiError) {
	var service = &Service{}
	err := service.init(keyPath, region, withDecryption, nil)
	if err != nil{
		return nil, err
	}

	return service, nil
}

//NewService this method will return a new instance of ParameterStoreService
func NewServiceMock(keyPath string, region string, withDecryption bool, ssmService ISsmService) (*Service, *error_utils.ApiError) {
	var service = &Service{}
	err := service.init(keyPath, region, withDecryption, ssmService)
	if err != nil{
		return nil, err
	}

	return service, nil
}

//init will initialize defaults
func (a *Service) init(keyPath string, region string, withDecryption bool, ssmService ISsmService) *error_utils.ApiError {

	if len(strings.TrimSpace(keyPath)) == 0 {
		return error_utils.NewBadRequestError("Parameter Store: Key path is empty")
	}

	if len(strings.TrimSpace(region)) == 0 {
		return error_utils.NewBadRequestError("Parameter Store: Region is empty")
	}

	// if it is not mocked load normal ssm
	if ssmService == nil {
		service, err := a.getSsm()
		if err != nil {
			return error_utils.NewBadRequestError(fmt.Sprintf("Cloud Storage: Uploader err %v", err))
		}
		a.ssmService = service
	}else{
		a.ssmService = ssmService
	}

	// assign params
	a.keyPath = keyPath
	a.region = region
	a.withDecryption = withDecryption

	return nil
}

//GetParametersByPath this method will load all parameters available by path
func (a *Service) GetParametersByPath() (map[string]interface{}, *error_utils.ApiError) {

	withDecryption := a.withDecryption
	keyPath := a.keyPath
	result, err := a.ssmService.GetParametersByPath(&ssm.GetParametersByPathInput{
		Path:           &keyPath,
		WithDecryption: &withDecryption,
	})

	// check error
	if err != nil {
		return nil, error_utils.NewInternalServerError(fmt.Sprintf("parameterStore: GetParametersByPath : An error had occurred in parameters by path - %v", err))
	}

	// init response
	resp := map[string]interface{}{}
	for _, param := range result.Parameters {

		paramName := *param.Name
		paramValue := *param.Value

		resp[paramName] = paramValue
	}

	return resp, nil
}

//getSsm will get ssm instance
func (a *Service) getSsm() (ISsmService, error){
	sess, err := session.NewSessionWithOptions(session.Options{
		Config:            aws.Config{Region: aws.String(a.region)},
		SharedConfigState: session.SharedConfigEnable,
	})

	// check error
	if err != nil {
		return nil, err
	}

	// get parameters by path
	ssm := ssm.New(sess, aws.NewConfig().WithRegion(a.region))

	return ssm, nil
}
