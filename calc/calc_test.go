package calc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContains(t *testing.T) {
	assert := assert.New(t)
	assert.True(contains("("))
}

func TestIsCorrectInput(t *testing.T) {
	assert := assert.New(t)
	str := "(10+3) *7"
	arr, err := IsCorrectInput(str)
	assert.Nil(err)
	assert.Equal(arr[2], "10")
	assert.Equal(arr[6], "*")
}

func TestIsEndTrue(t *testing.T) {
	assert := assert.New(t)
	assert.True(isEnd("start", "end"))
}

func TestIsEndFalse(t *testing.T) {
	assert := assert.New(t)
	assert.False(isEnd("/", "end"))
}

func TestIsErrorTrue(t *testing.T) {
	assert := assert.New(t)
	assert.True(isError(")", "("))
}

func TestIsErrorFalse(t *testing.T) {
	assert := assert.New(t)
	assert.False(isError("(", "("))
}

func TestIsOneBasicTrue(t *testing.T) {
	assert := assert.New(t)
	assert.True(isOneBasic("(", ")"))
}

func TestIsOneBasicFalse(t *testing.T) {
	assert := assert.New(t)
	assert.False(isOneBasic("(", "end"))
}

var isEndOfBaseTrueTests = []struct {
	prev    string
	current string
}{
	{"+", "-"},
	{"-", "-"},
	{"*", "-"},
	{"/", "+"},
	{")", ")"},
	{")", "end"},
}

func TestIsEndOfBaseTrue(t *testing.T) {
	assert := assert.New(t)

	for _, v := range isEndOfBaseTrueTests {
		assert.True(isEndOfBase(v.prev, v.current))
	}
}

var isEndOfBaseFalseTests = []struct {
	prev    string
	current string
}{
	{"start", "end"},
	{"(", ")"},
	{")", "("},
	{"start", "*"},
	{"(", "+"},
	{"*", "("},
}

func TestIsEndOfBaseFalse(t *testing.T) {
	assert := assert.New(t)

	for _, v := range isEndOfBaseFalseTests {
		assert.False(isEndOfBase(v.prev, v.current))
	}
}

var CalculateTrueTests = []struct {
	left      string
	operation string
	right     string
	result    string
}{
	{"5", "+", "3", "8"},
	{"5", "-", "3", "2"},
	{"5", "*", "3", "15"},
	{"9", "/", "3", "3"},
}

func TestCalculateTrue(t *testing.T) {
	assert := assert.New(t)

	for _, v := range CalculateTrueTests {
		result, err := calculate(v.left, v.operation, v.right)
		assert.Nil(err)
		assert.Equal(result, v.result)
	}
}

var CalculateFalseTests = []struct {
	left      string
	operation string
	right     string
}{
	{"5", "^", "3"},
	{"a", "/", "3"},
	{"5", "^", "v"},
	{"6", "/", "0"},
}

func TestCalculateFalse(t *testing.T) {
	assert := assert.New(t)

	for _, v := range CalculateFalseTests {
		_, err := calculate(v.left, v.operation, v.right)
		assert.Error(err)
	}
}

var CalcTrueTests = []struct {
	slice  []string
	result int
}{
	{[]string{"start", "1", "+", "3", "*", "5", "end"}, 16},
	{[]string{"start", "(", "1", "+", "3", ")", "*", "5", "end"}, 20},
	{[]string{"start", "(", "1", "+", "3", ")", "*", "(", "5", "-", "4", ")", "end"}, 4},
	{[]string{"start", "(", "1", ")", "+", "3", "*", "(", "5", "-", "4", ")", "end"}, 4},
}

func TestCalcTrue(t *testing.T) {
	assert := assert.New(t)
	for _, v := range CalcTrueTests {
		result, err := Calc(v.slice)
		assert.Nil(err)
		assert.Equal(result, v.result)
	}
}

var CalcFalseTests = []struct {
	slice []string
}{
	{[]string{"start", ")", "1", "+", "3", ")", "*", "5", "end"}},
	{[]string{"start", "(", "1", "+", "(", "3", ")", "*", "5", "end"}},
	{[]string{"start", "(", "1", "+", "(", "b", ")", "*", "5", "end"}},
}

func TestCalcFalse(t *testing.T) {
	assert := assert.New(t)
	for _, v := range CalcFalseTests {
		_, err := Calc(v.slice)
		assert.Error(err)
	}
}
