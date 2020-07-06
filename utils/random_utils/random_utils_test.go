package random_utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUUID(t *testing.T){

	//arrange
	expectedUuidLength := 36

	//act
	uuid, err := NewUUID()

	//assert
	assert.Nil(t, err)
	assert.NotNil(t, uuid)
	assert.EqualValues(t, expectedUuidLength, len(uuid))
}

func TestNewRandomStringLength(t *testing.T){

	//arrange
	expectedLength := 4

	//act
	result := NewRandomString(expectedLength)

	//assert
	assert.NotNil(t, result)
	assert.EqualValues(t, expectedLength, len(result))
}

func TestNewRandomStringNotTheSame(t *testing.T){
	//arrange
	expectedLength := 3

	//act
	resultOne := NewRandomString(expectedLength)
	resultTwo := NewRandomString(expectedLength)

	//assert
	assert.NotNil(t, resultOne)
	assert.NotNil(t, resultTwo)
	assert.NotEqual(t, resultOne, resultTwo)
	assert.EqualValues(t, expectedLength, len(resultOne))
	assert.EqualValues(t, expectedLength, len(resultTwo))
}
