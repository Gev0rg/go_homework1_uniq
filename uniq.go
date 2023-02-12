package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	// All Flags
	var numOccurrencesStrings bool
	flag.BoolVar(&numOccurrencesStrings, "c", false, "print the number of string occurrences")
	var onlyOccurrencesStrings bool
	flag.BoolVar(&onlyOccurrencesStrings, "d", false, "print only occurrences strings")
	var onlyNotOccurrencesStrings bool
	flag.BoolVar(&onlyNotOccurrencesStrings, "u", false, "print only not occurrences strings")
	var numFields int
	flag.IntVar(&numFields, "f", 0, "number of fields to skip from the string")
	var numChars int
	flag.IntVar(&numChars, "s", 0, "number of characters to skip from the string")
	var caseInsensitive bool
	flag.BoolVar(&caseInsensitive, "i", false, "case-insensitive")

	flag.Parse()


	// Check for having the file
	var in io.Reader
	if filename := flag.Arg(0); filename != "" {
		f, err := os.Open(filename)
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
	buf := bufio.NewScanner(in)

	// Create slice to store the lines
	strings := []string{}

	for ; buf.Scan(); {
		strings = append(strings, buf.Text()) 
	}

	// Check for errors from the scanner
	if err := buf.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading: err:", err)
	}

	for _, i := range(strings) {
		fmt.Println(i)
	}
}
