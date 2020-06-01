package cloudstorage

import (
	"bytes"
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/googleapis/google-cloud-go-testing/storage/stiface"
	"github.com/lelinu/api_utils/utils/error_utils"
	"google.golang.org/api/option"
	"io"
	"io/ioutil"
)

//Service struct
type Service struct {
	mock   bool
	client stiface.Client
	ctx    context.Context
}

// NewService initializes a gcp client service
func NewService(keystoreFilePath string, mock bool) (*Service, *error_utils.ApiError) {

	var service = &Service{}
	err := service.init(keystoreFilePath, mock)
	if err != nil {
		return nil, err
	}

	return service, nil
}

//init will initialize defaults
func (s *Service) init(keystoreFilePath string, mock bool) *error_utils.ApiError {

	if len(keystoreFilePath) == 0 {
		return error_utils.NewBadRequestError("Cloud Storage: KeystoreFilePath is empty")
	}

	// set context
	ctx := context.Background()

	if !mock{
		c, err := storage.NewClient(ctx, option.WithCredentialsFile(keystoreFilePath))
		if err != nil {
			return error_utils.NewInternalServerError(err.Error())
		}
		s.client = stiface.AdaptClient(c)
	}

	// set properties
	s.ctx = ctx

	return nil
}

//UploadFile will upload from file bytes
func (s *Service) UploadFile(fileName string, bucketName string, fileData []byte) *error_utils.ApiError {

	bh := s.client.Bucket(bucketName)
	_, err := bh.Attrs(s.ctx)
	if err != nil {
		return error_utils.NewInternalServerError(fmt.Sprintf("cloudStorage: UploadFile : Bucket error %v : %v", bucketName, err))
	}

	w := bh.Object(fileName).NewWriter(s.ctx)

	if _, err := io.Copy(w, bytes.NewReader(fileData)); err != nil {
		return error_utils.NewInternalServerError(fmt.Sprintf("cloudStorage: UploadFile : Unable to copy object %v : %v", fileName, err))
	}
	if err := w.Close(); err != nil {
		return error_utils.NewInternalServerError(fmt.Sprintf("cloudStorage: UploadFile : Unable to close writer %v : %v", fileName, err))
	}

	return nil
}

//DownloadFile will download a file from cloud storage
func (s *Service) DownloadFile(fileName string, bucketName string) ([]byte, *error_utils.ApiError) {

	bh := s.client.Bucket(bucketName)
	_, err := bh.Attrs(s.ctx)
	if err != nil {
		return nil, error_utils.NewInternalServerError(fmt.Sprintf("cloudStorage: DownloadFile : Bucket error %v : %v", bucketName, err))
	}

	rc, err := bh.Object(fileName).NewReader(s.ctx)
	if err != nil {
		return nil, error_utils.NewInternalServerError(fmt.Sprintf("cloudStorage: DownloadFile : Unable to find file %v : %v", fileName, err))
	}
	defer rc.Close()

	data, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, error_utils.NewInternalServerError(fmt.Sprintf("cloudStorage: DownloadFile : Unable to read bytes %v : %v", fileName, err))
	}
	return data, nil
}

//DeleteFile will delete a file from cloud storage
func (s *Service) DeleteFile(fileName string, bucketName string) *error_utils.ApiError {

	bh := s.client.Bucket(bucketName)
	_, err := bh.Attrs(s.ctx)
	if err != nil {
		return error_utils.NewInternalServerError(fmt.Sprintf("cloudStorage: DeleteFile : Bucket error %v : %v", bucketName, err))
	}

	o := bh.Object(fileName)

	if err := o.Delete(s.ctx); err != nil {
		return error_utils.NewInternalServerError(fmt.Sprintf("cloudStorage: DeleteFile : Unable to delete file %v : %v", fileName, err))
	}

	return nil
}
