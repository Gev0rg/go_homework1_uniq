package calc

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestContains (t *testing.T) {
	assert := assert.New(t)
	assert.True(contains("("))
}

func TestIsCorrectInput (t *testing.T) {
	assert := assert.New(t)
	str := "(10+3) *7"
	arr, err := IsCorrectInput(str)
	assert.Nil(err)
	assert.Equal(arr[2], "10")
	assert.Equal(arr[6], "*")
}
