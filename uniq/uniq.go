package uniq

import (
	"errors"
	"flag"
	"strconv"
	"strings"
)

// All Flags
var Options struct {
	numOccurrencesStrings     bool
	onlyOccurrencesStrings    bool
	onlyNotOccurrencesStrings bool
	numFields                 int
	numChars                  int
	caseInsensitive           bool
}

// Check all flags
func CheckFlags() error {
	// Init Flags
	flag.BoolVar(&Options.numOccurrencesStrings, "c", false, "print the number of string occurrences")
	flag.BoolVar(&Options.onlyOccurrencesStrings, "d", false, "print only occurrences inputStrings")
	flag.BoolVar(&Options.onlyNotOccurrencesStrings, "u", false, "print only not occurrences inputStrings")
	flag.IntVar(&Options.numFields, "f", 0, "number of fields to skip from the string")
	flag.IntVar(&Options.numChars, "s", 0, "number of characters to skip from the string")
	flag.BoolVar(&Options.caseInsensitive, "i", false, "case-insensitive")

	// Check for having the flags
	flag.Parse()

	if Options.onlyOccurrencesStrings && Options.onlyNotOccurrencesStrings {
		return errors.New("flags -d and -u are used simultaneously")
	}

	return nil
}

// Get unique strings from inputStrings
func GetUniqSlice(inputStrings []string, compareStrings []string) ([]string, []int) {
	outputStrings := []string{}
	numRepeatStrings := []int{}
	numRepeatStrings = append(numRepeatStrings, 0)

	indexNumRepeatStrings := 0
	prev := 0
	for i, s := range compareStrings {
		if s == compareStrings[prev] {
			numRepeatStrings[indexNumRepeatStrings]++
			continue
		}
		outputStrings = append(outputStrings, inputStrings[prev])
		numRepeatStrings = append(numRepeatStrings, 0)
		indexNumRepeatStrings++
		numRepeatStrings[indexNumRepeatStrings]++
		prev = i
	}
	outputStrings = append(outputStrings, inputStrings[prev])

	return outputStrings, numRepeatStrings
}

// Get output strings from inputStrings
func GetOutputSlice(inputStrings []string) []string {
	// Create slice to store the lines for compare after flag parsing
	compareStrings := make([]string, len(inputStrings))
	copy(compareStrings, inputStrings)

	// Flag parsing
	if Options.caseInsensitive {
		for i, s := range compareStrings {
			compareStrings[i] = strings.ToLower(s)
		}
	}

	if Options.numFields != 0 {
		for i, s := range compareStrings {
			arr := strings.Split(s, " ")[Options.numFields:]
			compareStrings[i] = strings.Join(arr, " ")
		}
	}

	if Options.numChars != 0 {
		for i, s := range compareStrings {
			if compareStrings[i] != "" {
				compareStrings[i] = s[Options.numChars:]
			}
		}
	}

	// Create slice to store the result lines
	compareStrings, numRepeatStrings := GetUniqSlice(inputStrings, compareStrings)

	if Options.numOccurrencesStrings {
		for i, s := range compareStrings {
			compareStrings[i] = strconv.Itoa(numRepeatStrings[i]) + " " + s
		}
	}

	outputStrings := []string{}

	switch true {
	case Options.onlyOccurrencesStrings:
		for i, s := range compareStrings {
			if numRepeatStrings[i] > 1 {
				outputStrings = append(outputStrings, s)
			}
		}
	case Options.onlyNotOccurrencesStrings:
		for i, s := range compareStrings {
			if numRepeatStrings[i] == 1 {
				outputStrings = append(outputStrings, s)
			}
		}
	default:
		outputStrings = compareStrings
	}

	return outputStrings
}
