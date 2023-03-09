package calc

import (
	"errors"
	"strconv"
	"strings"
)

type Stack struct {
	stack []string
}

func (s *Stack) Push(value string) {
	s.stack = append(s.stack, value)
}

func (s *Stack) Pop() string {
	top := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return top
}

func (s *Stack) Peek() string {
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

func getStringNumber(expression string, position *int) string {
	var strNum string
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

func toPostfix(infixExpression string) ([]string, error) {
	var stack Stack
	var postfixExpression []string
	splitInfixExpression := strings.Split(infixExpression, "")

	for position := 0; position < len(splitInfixExpression); position++ {
		cur := splitInfixExpression[position]
		_, isInt := strconv.Atoi(cur)
		_, isCurInOperation := operationPriority[cur]

		switch {
		case isInt == nil:
			postfixExpression = append(postfixExpression, getStringNumber(infixExpression, &position))
		case cur == "(":
			stack.Push(cur)
		case cur == ")":
			for !stack.isEmpty() && stack.Peek() != "(" {
				postfixExpression = append(postfixExpression, stack.Pop())
			}
			if stack.isEmpty() {
				return nil, errors.New("Invalid ')' in " + strconv.Itoa(position) + " position of expression")
			}
			stack.Pop()
		case isCurInOperation:
			if cur == "-" && (position == 0 || splitInfixExpression[position - len(postfixExpression[len(postfixExpression) - 1])] == "(") {
				cur = "~"
			}
			for !stack.isEmpty() && operationPriority[stack.Peek()] >= operationPriority[cur] {
				postfixExpression = append(postfixExpression, stack.Pop())
			}
			stack.Push(cur)
		}
	}

	for !stack.isEmpty() {
		postfixExpression = append(postfixExpression, stack.Pop())
	}

	return postfixExpression, nil
}

func execute(first string, operand string, second string) (string, error) {
	var result int

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

func calculatePostfix(postfixExpression []string) (string, error) {
	var stack Stack

	for position := 0; position < len(postfixExpression); position++ {
		cur := postfixExpression[position]

		if _, err := strconv.Atoi(cur); err == nil {
			stack.Push(cur)
		} else if _, isCurInOperation := operationPriority[cur]; isCurInOperation {
			// if unary '-'
			if cur == "~" {
				var top string
				if stack.isEmpty() {
					return "", errors.New("Invalid unary '-' in " + strconv.Itoa(position) + " position")
				}
				top = stack.Pop()

				exec, err := execute("0", "-", top)
				if err != nil {
					return "", err
				}
				stack.Push(exec)

				continue
			}

			if stack.isEmpty() {
				return "", errors.New("invalid " + cur + " in " + strconv.Itoa(position) + " position of postfix expression")
			}
			second := stack.Pop()
			if stack.isEmpty() {
				return "", errors.New("invalid " + cur + " in " + strconv.Itoa(position) + " position of postfix expression")
			}
			first := stack.Pop()

			exec, err := execute(first, cur, second)
			if err != nil {
				return "", err
			}
			stack.Push(exec)
		}
	}

	if stack.isEmpty() {
		return "", errors.New("result is not in stack")
	}
	return stack.Pop(), nil
}

func Run(infixExpression string) (string, error) {
	infixExpression = strings.ReplaceAll(infixExpression, " ", "")
	postfix, err := toPostfix(infixExpression)
	if err != nil {
		return "", err
	}

	result, err := calculatePostfix(postfix)
	if err != nil {
		return "", err
	}

	return result, nil
}
