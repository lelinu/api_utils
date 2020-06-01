package cloudstorage

import (
	"context"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

var (
	bucketName       = "dev-test"
	keyStoreFilePath = "keystore/service_account.json"
	fileName         = "Capture.JPG"
	filePath         = "files/Capture.JPG"
)

func TestNewServiceEmptyKeyStorePath(t *testing.T){
	// setup
	service, apiError := NewService("", true)

	// assert
	assert.NotNil(t, apiError)
	assert.Nil(t, service)
}

//TestUploadFileSuccess
func TestUploadFileSuccess(t *testing.T) {

	// read file
	dat, err := ioutil.ReadFile(filePath)
	assert.Nil(t, err)
	assert.NotNil(t, dat)

	// create a fake client
	client := newFakeClient()
	bkt := client.Bucket(bucketName)
	bkt.Create(context.Background(), bucketName, nil)

	// init service mock
	service, apiError := NewService(keyStoreFilePath, true)
	assert.Nil(t, apiError)
	assert.NotNil(t, service)

	service.client = client

	// upload file
	apiError = service.UploadFile(fileName, bucketName, dat)
	assert.Nil(t, apiError)
}

//TestUploadFileInvalidBucket
func TestUploadFileInvalidBucket(t *testing.T) {

	dat, err := ioutil.ReadFile(filePath)
	assert.Nil(t, err)
	assert.NotNil(t, dat)

	// init new fake client
	client := newFakeClient()

	// init service mock
	service, apiError := NewService(keyStoreFilePath, true)
	assert.Nil(t, apiError)
	assert.NotNil(t, service)

	service.client = client

	// upload file
	apiError = service.UploadFile(fileName, "doesnotexistbucket", dat)
	assert.NotNil(t, apiError)
	assert.EqualValues(t, "cloudStorage: UploadFile : Bucket error doesnotexistbucket : bucket \"doesnotexistbucket\" does not exist", apiError.ErrorMessage)
}

//TestDownloadFileNotFound
func TestDownloadFileNotFound(t *testing.T) {

	// init new fake client
	client := newFakeClient()
	bkt := client.Bucket(bucketName)
	bkt.Create(context.Background(), bucketName, nil)

	// init service mock
	service, apiError := NewService(keyStoreFilePath, true)
	assert.Nil(t, apiError)
	assert.NotNil(t, service)

	service.client = client

	// download file
	bytes, apiError := service.DownloadFile(fileName, bucketName)
	assert.NotNil(t, apiError)
	assert.EqualValues(t, "cloudStorage: DownloadFile : Unable to find file Capture.JPG : object \"Capture.JPG\" not found in bucket \"dev-test\"", apiError.ErrorMessage)
	assert.Nil(t, bytes)
}

//TestDownloadFileValid
func TestDownloadFileValid(t *testing.T) {

	// init new fake client
	client := newFakeClient()
	bkt := client.Bucket(bucketName)
	bkt.Create(context.Background(), bucketName, nil)

	// init service mock
	service, apiError := NewService(keyStoreFilePath, true)
	assert.Nil(t, apiError)
	assert.NotNil(t, service)

	service.client = client

	// read file
	dat, err := ioutil.ReadFile(filePath)
	assert.Nil(t, err)
	assert.NotNil(t, dat)

	// upload file
	apiError = service.UploadFile(fileName, bucketName, dat)
	assert.Nil(t, apiError)

	// download file
	bytes, apiError := service.DownloadFile(fileName, bucketName)
	assert.Nil(t, apiError)
	assert.NotNil(t, bytes)

	// compare length
	assert.Nil(t, err)
	assert.NotNil(t, dat)
	assert.EqualValues(t, len(bytes), len(dat))
}

//TestDownloadFileInvalidBucket
func TestDownloadFileInvalidBucket(t *testing.T) {

	// init new fake client
	client := newFakeClient()

	// init service mock
	service, apiError := NewService(keyStoreFilePath, true)
	assert.Nil(t, apiError)
	assert.NotNil(t, service)

	service.client = client

	// download file
	dat, err := service.DownloadFile(fileName, "doesnotexistbucket")
	assert.Nil(t, dat)
	assert.NotNil(t, err)
	assert.EqualValues(t, "cloudStorage: DownloadFile : Bucket error doesnotexistbucket : bucket \"doesnotexistbucket\" does not exist", err.ErrorMessage)
}

//TestDeleteFileSuccess will test the delete functionality
func TestDeleteFileSuccess(t *testing.T) {

	// init new fake client
	client := newFakeClient()
	bkt := client.Bucket(bucketName)
	bkt.Create(context.Background(), bucketName, nil)

	// init service mock
	service, apiError := NewService(keyStoreFilePath, true)
	assert.Nil(t, apiError)
	assert.NotNil(t, service)

	service.client = client

	// read file
	dat, err := ioutil.ReadFile(filePath)
	assert.Nil(t, err)
	assert.NotNil(t, dat)

	// upload file
	apiError = service.UploadFile(fileName, bucketName, dat)
	assert.Nil(t, apiError)

	// delete file
	apiError = service.DeleteFile(fileName, bucketName)
	assert.Nil(t, apiError)
}

//TestDownloadFileInvalidBucket will test the download file for an in-existing file
func TestDeleteFileInvalidBucket(t *testing.T) {

	// init new fake client
	client := newFakeClient()

	// init service mock
	service, apiError := NewService(keyStoreFilePath, true)
	assert.Nil(t, apiError)
	assert.NotNil(t, service)

	service.client = client

	// delete file
	apiError = service.DeleteFile(fileName, bucketName)
	assert.NotNil(t, apiError)
	assert.EqualValues(t, "cloudStorage: DeleteFile : Bucket error dev-test : bucket \"dev-test\" does not exist", apiError.ErrorMessage)
}
