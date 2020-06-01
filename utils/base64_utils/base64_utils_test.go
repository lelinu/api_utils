package base64_utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBase64EncodeValid(t *testing.T){
	//arrange
	input := "hello"

	//act
	val := EncodeFromString(input)
	fmt.Println(val)

	// assert
	assert.NotNil(t, val)
}

func TestEncodeFromBytesValid(t *testing.T){
	//arrange
	input := "hello"

	//act
	val := EncodeFromBytes([]byte(input))

	// assert
	assert.NotNil(t, val)
}

func TestDecodeToBytesValid(t *testing.T){
	//arrange
	input := "aGVsbG8="

	//act
	val, err := DecodeToBytes(input)

	// assert
	assert.Nil(t, err)
	assert.NotNil(t, val)
}

func TestDecodeToBytesInvalid(t *testing.T){
	//arrange
	input := "aGVsbG8"

	//act
	val, err := DecodeToBytes(input)

	// assert
	assert.NotNil(t, err)
	assert.Nil(t, val)
}

func TestDecodeToStringValid(t *testing.T){
	//arrange
	input := "aGVsbG8="

	//act
	val, err := DecodeToString(input)

	// assert
	assert.Nil(t, err)
	assert.NotNil(t, val)
}

func TestDecodeToStringInvalid(t *testing.T){
	//arrange
	input := "aGVsbG8"

	//act
	val, err := DecodeToString(input)

	// assert
	assert.NotNil(t, err)
	assert.EqualValues(t, "", val)
}



