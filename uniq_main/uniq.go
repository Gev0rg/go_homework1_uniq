package main

import (
	"github.com/Gev0rg/go_homework1_uniq/uniq"
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
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
	
		// Check for having the flags
		uniq.CheckFlags()
	
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
	
		// Get the output strings
		outputStrings := uniq.GetOutputSlice(inputStrings)
	
		for _, s := range outputStrings {
			fmt.Fprintln(out, s)
		}
}
