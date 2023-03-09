package calc

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetStringNumber(t *testing.T) {
	forGetNumberFromString := []struct {
		position   int
		expression string
		result     string
	}{
		{0, "156dkj68568", "156"},
	}

	for _, test := range forGetNumberFromString {
		result,_ := getStringNumber(test.expression, test.position)
		require.Equal(t, test.result, result, "Should get number from string")
	}
}

func TestToPostfix(t *testing.T) {
	forToPostfix := []struct {
		infix   string
		postfix []string
		err     error
	}{
		{"2+2+4567", []string{"2", "2", "+", "4567", "+"}, nil},
		{"10+4*5/6-3", []string{"10", "4", "5", "*", "6", "/", "+", "3", "-"}, nil},
		{"(2+2)*)8-13", nil, errors.New("")},
		{"(2+2)*8-13", []string{"2", "2", "+", "8", "*", "13", "-"}, nil},
	}

	for _, test := range forToPostfix {
		postfix, err := toPostfix(test.infix)
		if test.err != nil {
			require.Error(t, err)
		} else {
			require.Equal(t, test.postfix, postfix, "Should get postfix from infix")
		}
	}
}

func TestExecute(t *testing.T) {
	forToExecute := []struct {
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

	for _, test := range forToExecute {
		result, err := execute(test.first, test.operand, test.second)
		if test.err != nil {
			require.Error(t, err)
		} else {
			require.Equal(t, test.result, result, "Should get executed")
		}
	}
}

func TestCalculatePostfix(t *testing.T) {
	forCalculatePostfix := []struct {
		expression []string
		result     string
		err        error
	}{
		{[]string{"2", "2", "+"}, "4", nil},
		{[]string{"4", "0", "/"}, "", errors.New("")},
		{[]string{"2", "2", "+", "6", "*", "8", "2", "+", "-"}, "14", nil},
		{[]string{"2", "~", "2", "+"}, "0", nil},
	}

	for _, test := range forCalculatePostfix {
		result, err := calculatePostfix(test.expression)
		if test.err != nil {
			require.Error(t, err)
		} else {
			require.Equal(t, test.result, result, "Should get calculated postfix")
		}
	}
}

func TestRun(t *testing.T) {
	forRun := []struct {
		expression string
		result     string
		err        error
	}{
		{"2+2", "4", nil},
		{"4/0", "", errors.New("")},
		{"(2+2)*4-9/3", "13", nil},
		{"((2+2)*4-9/3", "", errors.New("")},
		{"19*13/13-13+(5-3)", "8", nil},
		{"19*13/13)-13+(5-3)", "", errors.New("")},
		{"16*2-(14*3-5)", "-5", nil},
		{"(5-9)*2", "-8", nil},
		{"(5-9)*2*(-1)", "8", nil},
		{"-(5-9)*2*(-1)", "-8", nil},
		{"19*13/13-13+(5-6)", "5", nil},
	}

	for _, test := range forRun {
		result, err := Run(test.expression)
		if test.err != nil {
			require.Error(t, err)
		} else {
			require.Equal(t, test.result, result, "Should get calculated postfix")
		}
	}
}
