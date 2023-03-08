package calc

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	stack []string
}

func (s *Stack) Push(value string) {
	s.stack = append(s.stack, value)
}

func (s *Stack) Pop() (string) {
	if len(s.stack) == 0 {
		fmt.Println("Error in Stack.Pop(): ", errors.New("stack is empty").Error())
		os.Exit(1)
	}

	top := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return top
}

func (s *Stack) Peek() (string) {
	if len(s.stack) == 0 {
		fmt.Println("Error in Stack.Peek(): ", errors.New("stack is empty").Error())
		os.Exit(1)
	}

	return s.stack[len(s.stack)-1]
}

func (s *Stack) isEmpty() bool {
	return len(s.stack) == 0
}

var operationPriority = map[string]int{
	"(": 1,
	"+": 2,
	"-": 2,
	"*": 3,
	"/": 3,
	"~": 4,
}

func getNumberFromString(expression string, position *int) string {
	strNum := ""
	splitExpression := strings.Split(expression, "")

	for ; *position < len(splitExpression); *position++ {
		num := splitExpression[*position]
		if _, err := strconv.Atoi(num); err == nil {
			strNum += num
		} else {
			*position--
			break
		}
	}

	return strNum
}

func toPostfix(infixExpression string) string {
	var stack Stack
	postfixExpression := ""
	splitInfixExpression := strings.Split(infixExpression, "")

	for position := range splitInfixExpression {
		cur := splitInfixExpression[position]

		if _, err:= strconv.Atoi(cur); err == nil {
			postfixExpression += getNumberFromString(infixExpression, &position)
		} else if cur == "(" {
			stack.Push(cur)
		} else if cur == ")" {
			for !stack.isEmpty() && stack.Peek() != "(" {
				postfixExpression += stack.Pop()
			}
			stack.Pop()			
		} else if _, haveCur := operationPriority[cur]; haveCur {
			if (cur == "-" && position == 0) {
				cur = "~"
			}
			for !stack.isEmpty() && operationPriority[stack.Peek()] >= operationPriority[cur] {
				postfixExpression += stack.Pop()
			}
			stack.Push(cur)
		}
	}

	for !stack.isEmpty() {
		postfixExpression += stack.Pop()
	}

	return postfixExpression
}

func execute(first string, operand string, second string) (string, error) {
	result := 0

	left, err := strconv.Atoi(first)
	if err != nil {
		return "", errors.New(first + " is not a number")
	}
	right, err := strconv.Atoi(second)
	if err != nil {
		return "", errors.New(second + " is not a number")
	}

	switch operand {
		case "+":
			result = left + right
		case "-":
			result = left - right
		case "*":
			result = left * right
		case "/":
			if right == 0 {
				return "", errors.New("division by zero")
			}
			result = left / right
	}

	return strconv.Itoa(result), nil
}

func calculatePostfix(postfixExpression string) (string, error) {
	var stack Stack
	splitPostfixExpression := strings.Split(postfixExpression, "")

	for position := range splitPostfixExpression {
		cur := splitPostfixExpression[position]

		if _, err:= strconv.Atoi(cur); err == nil {
			number := getNumberFromString(postfixExpression, &position)

			stack.Push(number)
		} else if _, haveCur := operationPriority[cur]; haveCur {
			if cur == "~" {
				var last string
				if !stack.isEmpty() {
					last = stack.Pop()
				} else {
					last = "0"
				}

				exec, err := execute("0", "-", last)
				if err != nil {
					return "", err
				}
				stack.Push(exec)

				continue
			}

			var second string
			if !stack.isEmpty() {
				second = stack.Pop()
			} else {
				second = "0"
			}
			var first string
			if !stack.isEmpty() {
				first = stack.Pop()
			} else {
				first = "0"
			}

			exec, err := execute(first, cur, second)
			if err != nil {
				return "", err
			}
			stack.Push(exec)
		}
	}

	return stack.Pop(), nil
}

func Run(infixExpression string) (string, error) {
	postfix := toPostfix(infixExpression)

	result, err := calculatePostfix(postfix)
	if err != nil {
		return "", err
	}

	return result, nil
}

// var runeContainer = [8]string{"(", ")", "+", "-", "*", "/", "start", "end"}

// // contains checks if a string is present in a slice
// func contains(str string) bool {
// 	for _, v := range runeContainer {
// 		if v == str {
// 			return true
// 		}
// 	}

// 	return false
// }

// func IsCorrectInput(str string) ([]string, error) {
// 	startArr := strings.Split(strings.ReplaceAll(str, " ", ""), "")
// 	finishArr := []string{"start"}
// 	prevIsNum := false
// 	for _, v := range startArr {
// 		if contains(v) {
// 			prevIsNum = false
// 			finishArr = append(finishArr, v)
// 			continue
// 		}
// 		if _, err := strconv.Atoi(v); err != nil {
// 			return nil, errors.New("invalid input " + v)
// 		}
// 		if prevIsNum {
// 			finishArr[len(finishArr)-1] += v
// 			continue
// 		}
// 		finishArr = append(finishArr, v)
// 		prevIsNum = true
// 	}
// 	finishArr = append(finishArr, "end")
// 	return finishArr, nil
// }

// func isEnd(prev string, current string) bool {
// 	return prev == "start" && current == "end"
// }

// func isError(prev string, current string) bool {
// 	return prev == "start" && current == ")" ||
// 		prev == "(" && current == "end" ||
// 		prev == ")" && current == "("
// }

// func isOneBasic(prev string, current string) bool {
// 	return prev == "(" && current == ")"
// }

// func isEndOfBase(prev string, current string) bool {
// 	return prev == "+" && current == "+" ||
// 		prev == "+" && current == "-" ||
// 		prev == "+" && current == ")" ||
// 		prev == "+" && current == "end" ||
// 		prev == "*" && current == "+" ||
// 		prev == "*" && current == "-" ||
// 		prev == "*" && current == "*" ||
// 		prev == "*" && current == "/" ||
// 		prev == "*" && current == ")" ||
// 		prev == "*" && current == "end" ||
// 		prev == ")" && current == "+" ||
// 		prev == ")" && current == "-" ||
// 		prev == ")" && current == "*" ||
// 		prev == ")" && current == "/" ||
// 		prev == ")" && current == ")" ||
// 		prev == ")" && current == "end" ||
// 		prev == "-" && current == "+" ||
// 		prev == "-" && current == "-" ||
// 		prev == "-" && current == ")" ||
// 		prev == "-" && current == "end" ||
// 		prev == "/" && current == "+" ||
// 		prev == "/" && current == "-" ||
// 		prev == "/" && current == "*" ||
// 		prev == "/" && current == "/" ||
// 		prev == "/" && current == ")" ||
// 		prev == "/" && current == "end"
// }

// func calculate(firstArg string, operation string, secondArg string) (string, error) {
// 	result := ""
// 	left, err := strconv.Atoi(firstArg)
// 	if err != nil {
// 		return result, err
// 	}
// 	right, err := strconv.Atoi(secondArg)
// 	if err != nil {
// 		return result, err
// 	}
// 	switch operation {
// 	case "+":
// 		result = fmt.Sprint(left + right)
// 	case "-":
// 		result = fmt.Sprint(left - right)
// 	case "*":
// 		result = fmt.Sprint(left * right)
// 	case "/":
// 		if right != 0 {
// 			result = fmt.Sprint(left / right)
// 		} else {
// 			err = errors.New("Invalid operation " + operation)
// 		}
// 	default:
// 		err = errors.New("Invalid operation " + operation)
// 	}

// 	return result, err
// }

// func Calc(arrFromStr []string) (int, error) {
// 	result := 0
// 	var err error = nil
// 	prev := arrFromStr[0]
// 	var stack stack.Stack
// 	stack.Push(prev)

// 	for i := 1; i < len(arrFromStr); i++ {
// 		if contains(arrFromStr[i]) {
// 			if isEnd(prev, arrFromStr[i]) {
// 				result, err = strconv.Atoi(fmt.Sprintf("%v", stack.Pop()))
// 				break
// 			}
// 			if isError(prev, arrFromStr[i]) {
// 				err = errors.New("invalid input " + arrFromStr[i])
// 				break
// 			}
// 			if isOneBasic(prev, arrFromStr[i]) {
// 				exp := stack.Pop()
// 				stack.Pop()
// 				prev = fmt.Sprintf("%v", stack.Peek())
// 				stack.Push(exp)
// 				continue
// 			}
// 			if isEndOfBase(prev, arrFromStr[i]) {
// 				secondArg := fmt.Sprintf("%v", stack.Pop())
// 				operation := fmt.Sprintf("%v", stack.Pop())
// 				firstArg := fmt.Sprintf("%v", stack.Pop())
// 				prev = fmt.Sprintf("%v", stack.Peek())

// 				result, err := calculate(firstArg, operation, secondArg)
// 				if err != nil {
// 					return 0, err
// 				}

// 				stack.Push(result)
// 				i--
// 				continue
// 			}
// 			prev = arrFromStr[i]
// 		}
// 		stack.Push(arrFromStr[i])
// 	}

// 	return result, err
// }
