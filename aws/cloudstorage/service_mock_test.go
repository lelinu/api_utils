package cloudstorage

import (
	"github.com/lelinu/api_utils/utils/error_utils"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

var (
	funcUploadFile func(fileName string, fileData []byte) (string, string, *error_utils.ApiError)
	funcDownloadFile func(fileKey string) ([]byte, *error_utils.ApiError)
)

type cloudStorageMock struct{}

func (a *cloudStorageMock) UploadFile(fileName string, fileData []byte) (string, string, *error_utils.ApiError) {
	return funcUploadFile(fileName, fileData)
}

func (a *cloudStorageMock) DownloadFile(fileKey string) ([]byte, *error_utils.ApiError) {
	return funcDownloadFile(fileKey)
}

//TestUploadFileMockValid will test the upload file to S3
func TestUploadFileMockValid(t *testing.T) {

	dat, err := ioutil.ReadFile("files/Capture.JPG")
	assert.Nil(t, err)
	assert.NotNil(t, dat)

	service := &cloudStorageMock{}
	funcUploadFile = func(fileName string, fileData []byte) (string, string, *error_utils.ApiError) {
		return "key", "location", nil
	}

	fileKey, objectURL, apiErr := service.UploadFile("Capture.JPG", dat)

	assert.Nil(t, apiErr)
	assert.NotNil(t, fileKey)
	assert.NotNil(t, objectURL)
}

//TestDownloadFileMockValid will test the download file
func TestDownloadFileMockValid(t *testing.T) {

	dat, err := ioutil.ReadFile("files/Capture.JPG")
	assert.Nil(t, err)
	assert.NotNil(t, dat)

	service := &cloudStorageMock{}
	funcDownloadFile = func(fileKey string) ([]byte, *error_utils.ApiError) {
		return dat, nil
	}

	bytes, apiErr := service.DownloadFile(folder + "/Capture.JPG")
	assert.Nil(t, apiErr)
	assert.NotNil(t, bytes)
}
