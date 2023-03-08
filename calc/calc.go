package calc

import (
	"errors"
	"fmt"
	"github.com/golang-collections/collections/stack"
	"strconv"
	"strings"
)

var runeContainer = [8]string{"(", ")", "+", "-", "*", "/", "start", "end"}

// contains checks if a string is present in a slice
func contains(str string) bool {
	for _, v := range runeContainer {
		if v == str {
			return true
		}
	}

	return false
}

func IsCorrectInput(str string) ([]string, error) {
	startArr := strings.Split(strings.ReplaceAll(str, " ", ""), "")
	finishArr := []string{"start"}
	prevIsNum := false
	for _, v := range startArr {
		if contains(v) {
			prevIsNum = false
			finishArr = append(finishArr, v)
			continue
		}
		if _, err := strconv.Atoi(v); err != nil {
			return nil, errors.New("invalid input " + v)
		}
		if prevIsNum {
			finishArr[len(finishArr)-1] += v
			continue
		}
		finishArr = append(finishArr, v)
		prevIsNum = true
	}
	finishArr = append(finishArr, "end")
	return finishArr, nil
}

func isEnd(prev string, current string) bool {
	return prev == "start" && current == "end"
}

func isError(prev string, current string) bool {
	return prev == "start" && current == ")" ||
		prev == "(" && current == "end" ||
		prev == ")" && current == "("
}

func isOneBasic(prev string, current string) bool {
	return prev == "(" && current == ")"
}

func isEndOfBase(prev string, current string) bool {
	return prev == "+" && current == "+" ||
		prev == "+" && current == "-" ||
		prev == "+" && current == ")" ||
		prev == "+" && current == "end" ||
		prev == "*" && current == "+" ||
		prev == "*" && current == "-" ||
		prev == "*" && current == "*" ||
		prev == "*" && current == "/" ||
		prev == "*" && current == ")" ||
		prev == "*" && current == "end" ||
		prev == ")" && current == "+" ||
		prev == ")" && current == "-" ||
		prev == ")" && current == "*" ||
		prev == ")" && current == "/" ||
		prev == ")" && current == ")" ||
		prev == ")" && current == "end" ||
		prev == "-" && current == "+" ||
		prev == "-" && current == "-" ||
		prev == "-" && current == ")" ||
		prev == "-" && current == "end" ||
		prev == "/" && current == "+" ||
		prev == "/" && current == "-" ||
		prev == "/" && current == "*" ||
		prev == "/" && current == "/" ||
		prev == "/" && current == ")" ||
		prev == "/" && current == "end"
}

func calculate(firstArg string, operation string, secondArg string) (string, error) {
	result := ""
	left, err := strconv.Atoi(firstArg)
	if err != nil {
		return result, err
	}
	right, err := strconv.Atoi(secondArg)
	if err != nil {
		return result, err
	}
	switch operation {
	case "+":
		result = fmt.Sprint(left + right)
	case "-":
		result = fmt.Sprint(left - right)
	case "*":
		result = fmt.Sprint(left * right)
	case "/":
		if right != 0 {
			result = fmt.Sprint(left / right)
		} else {
			err = errors.New("Invalid operation " + operation)
		}
	default:
		err = errors.New("Invalid operation " + operation)
	}

	return result, err
}

func Calc(arrFromStr []string) (int, error) {
	result := 0
	var err error = nil
	prev := arrFromStr[0]
	var stack stack.Stack
	stack.Push(prev)

	for i := 1; i < len(arrFromStr); i++ {
		if contains(arrFromStr[i]) {
			if isEnd(prev, arrFromStr[i]) {
				result, err = strconv.Atoi(fmt.Sprintf("%v", stack.Pop()))
				break
			}
			if isError(prev, arrFromStr[i]) {
				err = errors.New("invalid input " + arrFromStr[i])
				break
			}
			if isOneBasic(prev, arrFromStr[i]) {
				exp := stack.Pop()
				stack.Pop()
				prev = fmt.Sprintf("%v", stack.Peek())
				stack.Push(exp)
				continue
			}
			if isEndOfBase(prev, arrFromStr[i]) {
				secondArg := fmt.Sprintf("%v", stack.Pop())
				operation := fmt.Sprintf("%v", stack.Pop())
				firstArg := fmt.Sprintf("%v", stack.Pop())
				prev = fmt.Sprintf("%v", stack.Peek())

				result, err := calculate(firstArg, operation, secondArg)
				if err != nil {
					return 0, err
				}

				stack.Push(result)
				i--
				continue
			}
			prev = arrFromStr[i]
		}
		stack.Push(arrFromStr[i])
	}

	return result, err
}
