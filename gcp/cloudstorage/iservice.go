package cloudstorage

import "github.com/lelinu/api_utils/utils/error_utils"

//IService interface
type IService interface {
	UploadFile(fileName string, bucketName string, fileData []byte) *error_utils.ApiError
	DownloadFile(fileName string, bucketName string) ([]byte, *error_utils.ApiError)
	DeleteFile(fileName string, bucketName string) *error_utils.ApiError
}
