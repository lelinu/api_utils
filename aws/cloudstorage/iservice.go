package cloudstorage

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/lelinu/api_utils/utils/error_utils"
	"io"
)

//IService interface
type IService interface {
	UploadFile(fileName string, fileData []byte) (string, string, *error_utils.ApiError)
	DownloadFile(fileKey string) ([]byte, *error_utils.ApiError)
}

//IS3Uploader this is used generally for mocking
type IS3Uploader interface {
	Upload(input *s3manager.UploadInput, options ...func(*s3manager.Uploader)) (*s3manager.UploadOutput, error)
}

//IS3Downloader this is used generally for mocking
type IS3Downloader interface {
	Download(w io.WriterAt, input *s3.GetObjectInput, options ...func(*s3manager.Downloader)) (n int64, err error)
}