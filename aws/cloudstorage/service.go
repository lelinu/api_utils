package cloudstorage

import (
	"fmt"
	"github.com/lelinu/api_utils/utils/error_utils"
	"strings"

	"bytes"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

//Service struct
type Service struct {
	bucket     string
	folderName string
	region     string
	uploader   IS3Uploader
	downloader IS3Downloader
}

// NewService this method will return a new instance of CloudStorageService
func NewService(bucket string, folderName string, region string) (*Service, *error_utils.ApiError) {

	var service = &Service{}
	err := service.init(bucket, folderName, region, nil, nil)
	if err != nil{
		return nil, err
	}

	return service, nil
}

// NewServiceMock this service will return functionality for mocking purposes
func NewServiceMock(bucket string, folderName string, region string,
	uploader IS3Uploader, downloader IS3Downloader) (*Service, *error_utils.ApiError) {

	var service = &Service{}
	err := service.init(bucket, folderName, region, uploader, downloader)
	if err != nil{
		return nil, err
	}

	return service, nil
}

//UploadFile this method will upload file to s3 bucket
func (a *Service) UploadFile(fileName string, fileData []byte) (string, string, error) {

	// convert to reader
	r := bytes.NewReader(fileData)
	key := a.folderName + "/" + fileName

	// uploader to upload file
	output, err := a.uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(a.bucket),
		Key:    aws.String(key),
		Body:   r,
	})

	// give some information to the user if an error occurs
	if err != nil {
		return "", "", fmt.Errorf("cloudStorage: UploadFile :Unable to upload %v to %v, %v", fileName, a.bucket, err)
	}

	return key, output.Location, nil
}

//DownloadFile this method will download a file from s3 bucket
func (a *Service) DownloadFile(fileKey string) ([]byte, error) {

	// build request
	requestInput := &s3.GetObjectInput{
		Bucket: aws.String(a.bucket),
		Key:    aws.String(fileKey),
	}

	// init buffer
	buf := aws.NewWriteAtBuffer([]byte{})

	// download
	_, err := a.downloader.Download(buf, requestInput)
	if err != nil {
		return nil, fmt.Errorf("cloudStorage: DownloadFile : Unable to download %v from %v, %v", fileKey, a.bucket, err)
	}

	return buf.Bytes(), nil
}

//init will initialize defaults
func (a *Service) init(bucket string, folderName string, region string,
	uploader IS3Uploader, downloader IS3Downloader) *error_utils.ApiError {

	if len(strings.TrimSpace(bucket)) == 0 {
		return error_utils.NewBadRequestError("Cloud Storage: Bucket is empty.")
	}

	if len(strings.TrimSpace(folderName)) == 0 {
		return error_utils.NewBadRequestError("Cloud Storage: Folder name is empty.")
	}

	if len(strings.TrimSpace(region)) == 0 {
		return error_utils.NewBadRequestError("Cloud Storage: Region is empty.")
	}

	// if it is not mocked load normal s3 uploader
	if uploader == nil {
		uploader, err := a.getUploader()
		if err != nil {
			return error_utils.NewBadRequestError(fmt.Sprintf("Cloud Storage: Uploader err %v", err.Error()))
		}
		a.uploader = uploader
	}else{
		a.uploader = uploader
	}

	// if it is not mocked load normal s3 downloader
	if downloader == nil {
		downloader, err := a.getDownloader()
		if err != nil{
			return error_utils.NewBadRequestError(fmt.Sprintf("Cloud Storage: Downloader err %v", err.Error()))
		}
		a.downloader = downloader
	}else{
		a.downloader = downloader
	}

	// set properties
	a.bucket = bucket
	a.folderName = folderName
	a.region = region
	// set properties

	return nil
}

//getSession will get the aws session
func (a *Service) getSession() (*session.Session, error) {
	sess, err := session.NewSessionWithOptions(session.Options{
		Config:            aws.Config{Region: aws.String(a.region)},
		SharedConfigState: session.SharedConfigEnable,
	})

	return sess, err
}

//getUploader will get s3 uploader
func (a *Service) getUploader() (IS3Uploader, error){
	sess, err := a.getSession()
	if err != nil{
		return nil, err
	}

	uploader := s3manager.NewUploader(sess)
	return uploader, nil
}

//getDownloader will get s3 downloader
func (a *Service) getDownloader() (IS3Downloader, error){
	sess, err := a.getSession()
	if err != nil{
		return nil, err
	}

	downloader := s3manager.NewDownloader(sess)
	return downloader, nil
}