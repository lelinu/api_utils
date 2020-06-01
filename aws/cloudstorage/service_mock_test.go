package cloudstorage

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

var (
	funcUploadFile func(fileName string, fileData []byte) (string, string, error)
	funcDownloadFile func(fileKey string) ([]byte, error)
)

type cloudStorageMock struct{}

func (a *cloudStorageMock) UploadFile(fileName string, fileData []byte) (string, string, error) {
	return funcUploadFile(fileName, fileData)
}

func (a *cloudStorageMock) DownloadFile(fileKey string) ([]byte, error) {
	return funcDownloadFile(fileKey)
}

//TestUploadFileMockValid will test the upload file to S3
func TestUploadFileMockValid(t *testing.T) {

	dat, err := ioutil.ReadFile("files/Capture.JPG")
	assert.Nil(t, err)
	assert.NotNil(t, dat)

	service := &cloudStorageMock{}
	funcUploadFile = func(fileName string, fileData []byte) (string, string, error) {
		return "key", "location", nil
	}

	fileKey, objectURL, err := service.UploadFile("Capture.JPG", dat)

	assert.Nil(t, err)
	assert.NotNil(t, fileKey)
	assert.NotNil(t, objectURL)
}

//TestDownloadFileMockValid will test the download file
func TestDownloadFileMockValid(t *testing.T) {

	dat, err := ioutil.ReadFile("files/Capture.JPG")
	assert.Nil(t, err)
	assert.NotNil(t, dat)

	service := &cloudStorageMock{}
	funcDownloadFile = func(fileKey string) ([]byte, error) {
		return dat, nil
	}

	bytes, err := service.DownloadFile(folder + "/Capture.JPG")
	assert.Nil(t, err)
	assert.NotNil(t, bytes)
}
