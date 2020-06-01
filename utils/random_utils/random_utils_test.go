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
