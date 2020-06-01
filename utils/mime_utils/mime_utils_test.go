package mime_utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMimeTypeInvalidPath(t *testing.T) {

	// arrange
	p := "/hello/test"
	expected := "application/octet-stream"

	// act
	value := GetMimeType(p)

	// assert
	assert.NotNil(t, value)
	assert.EqualValues(t, expected, value)
}

func TestGetMimeTypeEmptyPath(t *testing.T) {

	// arrange
	p := ""
	expected := "application/octet-stream"

	// act
	value := GetMimeType(p)

	// assert
	assert.NotNil(t, value)
	assert.EqualValues(t, expected, value)
}

func TestGetMimeTypeValidPathJpg(t *testing.T) {

	// arrange
	p := "/hello/test.jpg"
	expected := "image/jpeg"

	// act
	value := GetMimeType(p)

	// assert
	assert.NotNil(t, value)
	assert.EqualValues(t, expected, value)
}

func TestGetMimeTypeValidPathJpeg(t *testing.T) {

	// arrange
	p := "/hello/test.JPEG"
	expected := "image/jpeg"

	// act
	value := GetMimeType(p)

	// assert
	assert.NotNil(t, value)
	assert.EqualValues(t, expected, value)
}

