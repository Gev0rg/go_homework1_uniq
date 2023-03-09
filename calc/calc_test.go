package calc

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

var forGetNumberFromString = []struct {
	position   int
	expression string
	result     string
}{
	{0, "156dkj68568", "156"},
}

func TestGetStringNumber(t *testing.T) {
	for _, test := range forGetNumberFromString {
		result := getStringNumber(test.expression, &test.position)
		require.Equal(t, test.result, result, "Should get number from string")
	}
}

var forToPostfix = []struct {
	infix   string
	postfix []string
	err     error
}{
	{"2+2+4567", []string{"2", "2", "+", "4567", "+"}, nil},
	{"10+4*5/6-3", []string{"10", "4", "5", "*", "6", "/", "+", "3", "-"}, nil},
	{"(2+2)*)8-13", nil, errors.New("")},
	{"(2+2)*8-13", []string{"2", "2", "+", "8", "*", "13", "-"}, nil},
}

func TestToPostfix(t *testing.T) {
	for _, test := range forToPostfix {
		postfix, err := toPostfix(test.infix)
		if test.err != nil {
			require.Error(t, err)
		} else {
			require.Equal(t, test.postfix, postfix, "Should get postfix from infix")
		}
	}
}

var forToExecute = []struct {
	first   string
	operand string
	second  string
	result  string
	err     error
}{
	{"5", "+", "2", "7", nil},
	{"t", "+", "2", "", errors.New("")},
	{"5", "+", "t", "", errors.New("")},
	{"7", "-", "3", "4", nil},
	{"4", "*", "2", "8", nil},
	{"9", "/", "3", "3", nil},
	{"4", "/", "0", "", errors.New("")},
}

func TestExecute(t *testing.T) {
	for _, test := range forToExecute {
		result, err := execute(test.first, test.operand, test.second)
		if test.err != nil {
			require.Error(t, err)
		} else {
			require.Equal(t, test.result, result, "Should get executed")
		}
	}
}

var forCalculatePostfix = []struct {
	expression []string
	result     string
	err        error
}{
	{[]string{"2", "2", "+"}, "4", nil},
	{[]string{"4", "0", "/"}, "", errors.New("")},
	{[]string{"2", "2", "+", "6", "*", "8", "2", "+", "-"}, "14", nil},
	{[]string{"2", "~", "2", "+"}, "0", nil},
}

func TestCalculatePostfix(t *testing.T) {
	for _, test := range forCalculatePostfix {
		result, err := calculatePostfix(test.expression)
		if test.err != nil {
			require.Error(t, err)
		} else {
			require.Equal(t, test.result, result, "Should get calculated postfix")
		}
	}
}

var forRun = []struct {
	expression string
	result     string
	err        error
}{
	{"2+2", "4", nil},
	{"4/0", "", errors.New("")},
}

func TestRun(t *testing.T) {
	for _, test := range forRun {
		result, err := Run(test.expression)
		if test.err != nil {
			require.Error(t, err)
		} else {
			require.Equal(t, test.result, result, "Should get calculated postfix")
		}
	}
}
