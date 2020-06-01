package env_utils

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetEnvWithoutFallbackSuccessful(t *testing.T){

	//arrange
	var key = "my-key"
	var value = "my-value"

	os.Setenv(key, value)

	//act
	retrievedValue := GetEnv("my-key", "")

	//assert
	assert.NotNil(t, retrievedValue)
	assert.EqualValues(t, value, retrievedValue)
}

func TestGetEnvWithFallbackSuccessful(t *testing.T){

	//arrange
	var fallback = "fallback"

	//act
	retrievedValue := GetEnv("my-key-2", fallback)

	//assert
	assert.NotNil(t, retrievedValue)
	assert.EqualValues(t, fallback, retrievedValue)
}