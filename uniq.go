package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// All Flags
	var numOccurrencesStrings bool
	flag.BoolVar(&numOccurrencesStrings, "c", false, "print the number of string occurrences")
	var onlyOccurrencesStrings bool
	flag.BoolVar(&onlyOccurrencesStrings, "d", false, "print only occurrences inputStrings")
	var onlyNotOccurrencesStrings bool
	flag.BoolVar(&onlyNotOccurrencesStrings, "u", false, "print only not occurrences inputStrings")
	var numFields int
	flag.IntVar(&numFields, "f", 0, "number of fields to skip from the string")
	var numChars int
	flag.IntVar(&numChars, "s", 0, "number of characters to skip from the string")
	var caseInsensitive bool
	flag.BoolVar(&caseInsensitive, "i", false, "case-insensitive")

	// Check for having the flags
	flag.Parse()

	// Check for having the input file
	var in io.Reader
	if inputFile := flag.Arg(0); inputFile != "" {
		f, err := os.Open(inputFile)
		if err != nil {
			fmt.Println("error opening file: err:", err)
			os.Exit(1)
		}
		defer f.Close()

		in = f
	} else {
		in = os.Stdin
	}

	// Create the scanner
	inBuf := bufio.NewScanner(in)

	// Create slice to store the input lines
	inputStrings := []string{}

	for ; inBuf.Scan(); {
		inputStrings = append(inputStrings, inBuf.Text()) 
	}

	// Check for errors from the scanner
	if err := inBuf.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading: err:", err)
	}

	// Create slice to store the result lines
	outputStrings := []string{}

	// Check for having the output file
	var out io.Writer
	if outputFile := flag.Arg(1); outputFile != "" {
		f, err := os.Create(outputFile)
		if err != nil {
			fmt.Println("error opening file: err:", err)
			os.Exit(1)
		}
		defer f.Close()

		out = f
	} else {
		out = os.Stdout
    }

	if caseInsensitive {
		var prev string
		for _, s := range inputStrings {
			if strings.ToLower(s) == strings.ToLower(prev) {
                continue
            }
			outputStrings = append(outputStrings, prev)
            prev = s
		}
		outputStrings = append(outputStrings, prev)
	}

	for _, s := range outputStrings[1:] {
		fmt.Fprintln(out, s)
	}
}
