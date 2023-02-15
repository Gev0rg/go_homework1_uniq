package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func getOutputSlice(inputStrings []string, compareStrings []string) ([]string, []int) {
	outputStrings := []string{}
	numStrings := []int{}
	numStrings = append(numStrings, 0)

	indexNumStrings := 0
	prev := 0
	for i, s := range compareStrings {
		if s == compareStrings[prev] {
			numStrings[indexNumStrings]++
			continue
		}
		outputStrings = append(outputStrings, inputStrings[prev])
		numStrings = append(numStrings, 0)
		indexNumStrings++
		numStrings[indexNumStrings]++
		prev = i
	}
	outputStrings = append(outputStrings, inputStrings[prev])

	return outputStrings, numStrings
}

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

	// Create slice to store the lines for compare after flag parsing
	compareStrings := make([]string, len(inputStrings))
	copy(compareStrings, inputStrings)

	// Flag parsing
	if caseInsensitive {
		for i, s := range compareStrings {
			compareStrings[i] = strings.ToLower(s)
		}
	}

	if numFields != 0 {
		for i, s := range compareStrings {
            arr := strings.Split(s, " ")[numFields:]
			compareStrings[i] = strings.Join(arr, " ")
        }
	}

	if numChars != 0 {
		for i, s := range compareStrings {
			if compareStrings[i] != "" {
				compareStrings[i] = s[numChars:]
			}
        }
    }

	// Create slice to store the result lines
	outputStrings, numStrings := getOutputSlice(inputStrings, compareStrings)

	if numOccurrencesStrings {
		for i, s := range outputStrings {
			outputStrings[i] = strconv.Itoa(numStrings[i]) + " " + s
		}
    }

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

    if onlyOccurrencesStrings && !onlyNotOccurrencesStrings {
		for i, s := range outputStrings {
			if numStrings[i] > 1 {
				fmt.Fprintln(out, s)
			}
		}
	} else if !onlyOccurrencesStrings && onlyNotOccurrencesStrings {
		for i, s := range outputStrings {
			if numStrings[i] == 1 {
				fmt.Fprintln(out, s)
			}
		}		
	} else {
		for _, s := range outputStrings {
			fmt.Fprintln(out, s)
		}
	}
}
