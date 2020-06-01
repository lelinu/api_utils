package cloudstorage

import (
	"errors"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	awsRegion = "eu-west-1"
	bucket    = "dev-test"
	folder    = "uploads"
)

var (
	funcS3Upload func(input *s3manager.UploadInput, options ...func(*s3manager.Uploader)) (*s3manager.UploadOutput, error)
	funcS3Download func(w io.WriterAt, input *s3.GetObjectInput, options ...func(*s3manager.Downloader)) (n int64, err error)
)

type s3UploaderMock struct{}
type s3DownloaderMock struct{}

func (a s3UploaderMock) Upload(input *s3manager.UploadInput, options ...func(*s3manager.Uploader)) (*s3manager.UploadOutput, error) {
	return funcS3Upload(input, options...)
}

func (a s3DownloaderMock) Download(w io.WriterAt, input *s3.GetObjectInput, options ...func(*s3manager.Downloader)) (n int64, err error){
	return funcS3Download(w, input, options...)
}

//TestUploadFileValid
func TestUploadFileValid(t *testing.T) {

	dat, err := ioutil.ReadFile("files/Capture.JPG")
	assert.Nil(t, err)
	assert.NotNil(t, dat)

	uploader := s3UploaderMock{}
	funcS3Upload = func(input *s3manager.UploadInput, options ...func(*s3manager.Uploader)) (*s3manager.UploadOutput, error){
		return &s3manager.UploadOutput{
			Location:  "Malta",
			VersionID: nil,
			UploadID:  "123456",
		}, nil
	}

	service, apiErr := NewServiceMock(bucket, folder, awsRegion, uploader, nil)
	assert.Nil(t, apiErr)
	assert.NotNil(t, service)

	fileKey, objectURL, err := service.UploadFile("Capture.JPG", dat)

	assert.Nil(t, err)
	assert.NotNil(t, fileKey)
	assert.EqualValues(t, "uploads/Capture.JPG", fileKey)
	assert.NotNil(t, objectURL)
	assert.EqualValues(t, "Malta", objectURL)
}

//TestUploadFileInvalid
func TestUploadFileInvalid(t *testing.T) {

	dat, err := ioutil.ReadFile("files/Capture.JPG")
	assert.Nil(t, err)
	assert.NotNil(t, dat)

	uploader := s3UploaderMock{}
	funcS3Upload = func(input *s3manager.UploadInput, options ...func(*s3manager.Uploader)) (*s3manager.UploadOutput, error){
		return &s3manager.UploadOutput{}, errors.New("invalid bucket name")
	}

	service, apiErr := NewServiceMock(bucket, folder, awsRegion, uploader, nil)
	assert.Nil(t, apiErr)
	assert.NotNil(t, service)

	fileKey, objectURL, err := service.UploadFile("Capture.JPG", dat)

	assert.NotNil(t, err)
	assert.EqualValues(t, "cloudStorage: UploadFile :Unable to upload Capture.JPG to dev-test, invalid bucket name", err.Error())
	assert.EqualValues(t, "", fileKey)
	assert.EqualValues(t, "", objectURL)
}

//TestDownloadFileValid
func TestDownloadFileValid(t *testing.T) {

	downloader := s3DownloaderMock{}
	funcS3Download = func(w io.WriterAt, input *s3.GetObjectInput, options ...func(*s3manager.Downloader)) (n int64, err error){
		return 1, nil
	}

	service, apiErr := NewServiceMock(bucket, folder, awsRegion, nil, downloader)
	assert.Nil(t, apiErr)
	assert.NotNil(t, service)

	bytes, err := service.DownloadFile("Capture.JPG")

	assert.Nil(t, err)
	assert.NotNil(t, bytes)
}

//TestDownloadFileInvalid
func TestDownloadFileInvalid(t *testing.T) {

	downloader := s3DownloaderMock{}
	funcS3Download = func(w io.WriterAt, input *s3.GetObjectInput, options ...func(*s3manager.Downloader)) (n int64, err error){
		return 0, errors.New("file not found")
	}

	service, apiErr := NewServiceMock(bucket, folder, awsRegion, nil, downloader)
	assert.Nil(t, apiErr)
	assert.NotNil(t, service)

	bytes, err := service.DownloadFile("Capture.JPG")

	assert.NotNil(t, err)
	assert.EqualValues(t, "cloudStorage: DownloadFile : Unable to download Capture.JPG from dev-test, file not found", err.Error())
	assert.Nil(t, bytes)
}

