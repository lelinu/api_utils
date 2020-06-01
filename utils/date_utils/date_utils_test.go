package date_utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestApiDateTimeFormatConstSuccessful(t *testing.T){
	assert.EqualValues(t, time.RFC3339, ApiDateTimeFormat)
}

func TestDbDateTimeFormatConstSuccessful(t *testing.T){
	assert.EqualValues(t, "2006-01-02 15:04:05", DbDateTimeFormat)
}

func TestGetApiCurrentDateTimeStringSuccessful(t *testing.T){

	// arrange
	expected := time.Now().UTC().Format(time.RFC3339)

	// act
	now := GetApiCurrentDateTimeString()

	// assert
	assert.NotNil(t, now)
	assert.EqualValues(t, expected, now)
}

func TestGetDbCurrentDateTimeString(t *testing.T){

	// arrange
	expected := time.Now().UTC().Format("2006-01-02 15:04:05")

	// act
	now := GetDbCurrentDateTimeString()

	// assert
	assert.NotNil(t, now)
	assert.EqualValues(t, expected, now)
}