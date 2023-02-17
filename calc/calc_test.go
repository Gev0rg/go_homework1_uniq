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

func TestIsEndTrue (t *testing.T) {
	assert := assert.New(t)
    assert.True(isEnd("start", "end"))
}

func TestIsEndFalse (t *testing.T) {
	assert := assert.New(t)
    assert.False(isEnd("/", "end"))
}

func TestIsErrorTrue (t *testing.T) {
	assert := assert.New(t)
    assert.True(isError(")", "("))
}

func TestIsErrorFalse (t *testing.T) {
	assert := assert.New(t)
    assert.False(isError("(", "("))
}

func TestIsOneBasicTrue (t *testing.T) {
	assert := assert.New(t)
    assert.True(isOneBasic("(", ")"))
}

func TestIsOneBasicFalse (t *testing.T) {
	assert := assert.New(t)
    assert.False(isOneBasic("(", "end"))
}

func TestIsEndOfBaseTrue (t *testing.T) {
	assert := assert.New(t)
    assert.True(isEndOfBase("+", "-"))
	assert.True(isEndOfBase("-", "-"))
	assert.True(isEndOfBase("*", "-"))
	assert.True(isEndOfBase("/", "+"))
	assert.True(isEndOfBase(")", ")"))
	assert.True(isEndOfBase(")", "end"))
}

func TestIsEndOfBaseFalse (t *testing.T) {
	assert := assert.New(t)
    assert.False(isEndOfBase("start", "end"))
	assert.False(isEndOfBase("(", ")"))
	assert.False(isEndOfBase(")", "("))
	assert.False(isEndOfBase("start", "*"))
	assert.False(isEndOfBase("(", "+"))
	assert.False(isEndOfBase("*", "("))
}

func TestCalculatePlus(t *testing.T) {
	assert := assert.New(t)
    result, err := calculate("5", "+", "3")
    assert.Nil(err)
    assert.Equal(result, "8")
}

func TestCalculateSub(t *testing.T) {
	assert := assert.New(t)
    result, err := calculate("5", "-", "3")
    assert.Nil(err)
    assert.Equal(result, "2")
}
func TestCalculateProd(t *testing.T) {
	assert := assert.New(t)
    result, err := calculate("5", "*", "3")
    assert.Nil(err)
    assert.Equal(result, "15")
}
func TestCalculateQuot(t *testing.T) {
	assert := assert.New(t)
    result, err := calculate("9", "/", "3")
    assert.Nil(err)
    assert.Equal(result, "3")
}

func TestCalculateErr1(t *testing.T) {
	assert := assert.New(t)
    _, err := calculate("5", "^", "3")
    assert.Error(err)
}

func TestCalculateErr2(t *testing.T) {
	assert := assert.New(t)
    _, err := calculate("a", "/", "3")
    assert.Error(err)
}

func TestCalculateErr3(t *testing.T) {
	assert := assert.New(t)
    _, err := calculate("5", "^", "v")
    assert.Error(err)
}

func TestCalcRight1(t *testing.T) {
    assert := assert.New(t)
	arrFromStr := []string{"start", "1", "+", "3", "*", "5", "end"}
	result, err := Calc(arrFromStr)
    assert.Nil(err)
	assert.Equal(result, 16)
}

func TestCalcRight2(t *testing.T) {
    assert := assert.New(t)
	arrFromStr := []string{"start", "(", "1", "+", "3", ")", "*", "5", "end"}
	result, err := Calc(arrFromStr)
    assert.Nil(err)
	assert.Equal(result, 20)
}

func TestCalcRight3(t *testing.T) {
    assert := assert.New(t)
	arrFromStr := []string{"start", "(", "1", "+", "3", ")", "*", "(", "5", "-", "4", ")", "end"}
	result, err := Calc(arrFromStr)
    assert.Nil(err)
	assert.Equal(result, 4)
}

func TestCalcRight4(t *testing.T) {
    assert := assert.New(t)
	arrFromStr := []string{"start", "(", "1", ")", "+", "3", "*", "(", "5", "-", "4", ")", "end"}
	result, err := Calc(arrFromStr)
    assert.Nil(err)
	assert.Equal(result, 4)
}

func TestCalcErr1(t *testing.T) {
    assert := assert.New(t)
	arrFromStr := []string{"start", ")", "1", "+", "3", ")", "*", "5", "end"}
	_, err := Calc(arrFromStr)
    assert.Error(err)
}

func TestCalcErr2(t *testing.T) {
    assert := assert.New(t)
	arrFromStr := []string{"start", "(", "1", "+", "(", "3", ")", "*", "5", "end"}
	_, err := Calc(arrFromStr)
    assert.Error(err)
}

func TestCalcErr3(t *testing.T) {
    assert := assert.New(t)
	arrFromStr := []string{"start", "(", "1", "+", "(", "b", ")", "*", "5", "end"}
	_, err := Calc(arrFromStr)
    assert.Error(err)
}
