package lzap

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/pkg/errors"
)

var (
	logger IService
)

func init() {
	logger = NewService("info", "access.log")
}

func getCurrentDirectory() string {
	ex, _ := os.Executable()
	return filepath.Dir(ex)
}

//TestPrintfWithTagsSuccessful will test the logger
func TestPrintfWithTagsSuccessful(t *testing.T) {
	logger.Printf("This is a message", "client:123456")
}

//TestPrintfWithoutTagsSuccessful will test the logger
func TestPrintfWithoutTagsSuccessful(t *testing.T) {
	logger.Printf("This is a message")
}

//TestLogInfoWithTagsSuccessful will test the logger
func TestLogInfoWithTagsSuccessful(t *testing.T) {
	logger.Info("This is a message", "client:123456")
}

//TestLogInfoWithoutTagsSuccessful will test the logger
func TestLogInfoWithoutTagsSuccessful(t *testing.T) {
	logger.Info("This is a message")
}

//TestLogWarnWithTagsSuccessful will test the logger
func TestLogWarnWithTagsSuccessful(t *testing.T) {
	logger.Warn("This is a warning message", "client:123456")
}

//TestLogWarnWithTagsSuccessful will test the logger
func TestLogWarnWithoutTagsSuccessful(t *testing.T) {
	logger.Warn("This is a warning message")
}

//TestLogDebugWithTagsSuccessful will test the logger
func TestLogDebugWithTagsSuccessful(t *testing.T) {
	logger.Debug("This is a debug message", "client:123456")
}

//TestLogDebugWithoutTagsSuccessful will test the logger
func TestLogDebugWithoutTagsSuccessful(t *testing.T) {
	logger.Debug("This is a debug message")
}

//TestLogErrorWithTagsSuccessful will test the logger
func TestLogErrorWithTagsSuccessful(t *testing.T) {
	logger.Error("This is an error message", errors.New("oops an error message"), "client:123456")
}

//TestLogErrorWithoutTagsSuccessful will test the logger
func TestLogErrorWithoutTagsSuccessful(t *testing.T) {
	logger.Error("This is an error message", errors.New("oops an error message"))
}

//BenchmarkLogInfo will benchmark the functionality
func BenchmarkLogInfo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		logger.Info("This is a message", "client:123456")
	}
}
