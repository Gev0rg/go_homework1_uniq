package uniq

import (
	"flag"
	"strconv"
	"strings"
)

// All Flags
var numOccurrencesStrings bool
var onlyOccurrencesStrings bool
var onlyNotOccurrencesStrings bool
var numFields int
var numChars int
var caseInsensitive bool

// Check all flags
func CheckFlags() {
	// Init Flags
	flag.BoolVar(&numOccurrencesStrings, "c", false, "print the number of string occurrences")
	flag.BoolVar(&onlyOccurrencesStrings, "d", false, "print only occurrences inputStrings")
	flag.BoolVar(&onlyNotOccurrencesStrings, "u", false, "print only not occurrences inputStrings")
	flag.IntVar(&numFields, "f", 0, "number of fields to skip from the string")
	flag.IntVar(&numChars, "s", 0, "number of characters to skip from the string")
	flag.BoolVar(&caseInsensitive, "i", false, "case-insensitive")

	// Check for having the flags
	flag.Parse()
}

// Get unique strings from inputStrings
func GetUniqSlice (inputStrings []string, compareStrings []string) ([]string, []int) {
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

// Get output strings from inputStrings
func GetOutputSlice(inputStrings[]string) []string {
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
	compareStrings, numStrings := GetUniqSlice(inputStrings, compareStrings)

	if numOccurrencesStrings {
		for i, s := range compareStrings {
			compareStrings[i] = strconv.Itoa(numStrings[i]) + " " + s
		}
    }

	outputStrings := []string{}
    if onlyOccurrencesStrings && !onlyNotOccurrencesStrings {
		for i, s := range compareStrings {
			if numStrings[i] > 1 {
				outputStrings = append(outputStrings, s)
			}
		}
	} else if !onlyOccurrencesStrings && onlyNotOccurrencesStrings {
		for i, s := range compareStrings {
			if numStrings[i] == 1 {
				outputStrings = append(outputStrings, s)
			}
		}		
	} else {
		outputStrings = compareStrings
	}

	return outputStrings
}
