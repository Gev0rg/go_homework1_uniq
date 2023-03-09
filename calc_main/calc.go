package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/Gev0rg/go_homework1_uniq/calc"
)

func main() {
	// Create the scanner
	inBuf := bufio.NewScanner(os.Stdin)
	inBuf.Scan()
	// Create string to store the input lines
	str := inBuf.Text()
	if str == "" {
		fmt.Println(errors.New("expression is empty").Error())
		os.Exit(1)
	}

	result, err := calc.Run(str)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(result)
}
