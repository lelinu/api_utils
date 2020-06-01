package crypto_utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateHashFromStringSuccessful(t *testing.T){

	//arrange
	var input = "my-input"
	var expectedHashLength = 60

	//act
	hashedInput, err := GenerateHashFromString(input)

	//assert
	assert.Nil(t, err)
	assert.NotNil(t, hashedInput)
	assert.EqualValues(t, expectedHashLength, len(hashedInput))
}

func TestGenerateHashAndSaltKeyFromInputSuccessful(t *testing.T){

	//arrange
	var input = "my-input"
	var expectedHashLength = 60
	var expectedSaltKeyLength = 36

	//act
	hashedInput, saltKey, err := GenerateHashAndSaltKeyFromInput(input)

	//assert
	assert.Nil(t, err)
	assert.NotNil(t, hashedInput)
	assert.EqualValues(t, expectedHashLength, len(hashedInput))
	assert.EqualValues(t, expectedSaltKeyLength, len(saltKey))
}

func TestCompareHashedPasswordSuccessful(t *testing.T){
	//arrange
	var input = "my-input"
	var expectedHashLength = 60
	var expectedSaltKeyLength = 36

	//act
	hashedInput, saltKey, err := GenerateHashAndSaltKeyFromInput(input)

	//assert
	assert.Nil(t, err)
	assert.NotNil(t, hashedInput)
	assert.EqualValues(t, expectedHashLength, len(hashedInput))
	assert.EqualValues(t, expectedSaltKeyLength, len(saltKey))

	err = CompareHashWithClearPasswordAndSalt(hashedInput, input, saltKey)
	assert.Nil(t, err)
}