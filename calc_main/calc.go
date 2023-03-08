package main

import (
	"bufio"
	"fmt"
	"github.com/Gev0rg/go_homework1_uniq/calc"
	"os"
)

func main() {
	// Create the scanner
	inBuf := bufio.NewScanner(os.Stdin)
	inBuf.Scan()
	// Create string to store the input lines
	str := inBuf.Text()
	// Check right input
	arrFromStr, err := calc.IsCorrectInput(str)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	// Calculate the input expression
	result, err := calc.Calc(arrFromStr)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(result)
}
