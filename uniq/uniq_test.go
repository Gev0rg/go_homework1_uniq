package uniq

import (
    "testing"
	"github.com/stretchr/testify/assert"
)

func TestGetUniqSlice1(t *testing.T) {
	assert := assert.New(t)
    slice := []string{"Father", "Father", "Mother", "Son", "Son", "Son"}
    uniq, _ := GetUniqSlice(slice, slice)
    assert.Equal(len(uniq), 3)
    assert.Equal(uniq[1], "Mother")
}

func TestGetOutputSlice1(t *testing.T) {
    assert := assert.New(t)
    slice := []string{"Father", "Father", "Mother", "Son", "Son", "Son"}
    uniq := GetOutputSlice(slice)
    assert.Equal(len(uniq), 3)
}

func TestGetOutputSlice2(t *testing.T) {
    assert := assert.New(t)

    numOccurrencesStrings = true
    onlyOccurrencesStrings = true
    numFields = 1

    slice := []string{"My Father", "Your Father", "Our Mother", "My Son", "Your Son", "Our Son"}

    uniq := GetOutputSlice(slice)
    assert.Equal(len(uniq), 2)
    assert.Equal(uniq[0], "2 My Father")
    assert.Equal(uniq[1], "3 My Son")

    numOccurrencesStrings = false
    onlyOccurrencesStrings = false
    numFields = 0
}

func TestGetOutputSlice3(t *testing.T) {
    assert := assert.New(t)

    onlyNotOccurrencesStrings = true
    numChars = 1
    caseInsensitive = true

    slice := []string{"M Father", "Y father", "O Mother", "M Son", "T Dad", "y SON", "O sOn"}

    uniq := GetOutputSlice(slice)
    assert.Equal(len(uniq), 3)
    assert.Equal(uniq[0], "O Mother")
    assert.Equal(uniq[1], "M Son")
    assert.Equal(uniq[2], "T Dad")

    onlyNotOccurrencesStrings = false
    numChars = 0
    caseInsensitive = false
}

func TestGetOutputSlice4(t *testing.T) {
    assert := assert.New(t)

    numFields = 1
    numChars = 2
    caseInsensitive = true
    numOccurrencesStrings = true

    slice := []string{"Pr M Father", "MR Y father", "FR O Mother", "NR M Son", "DR T Dad", "XR y SON", "Sr O sOn"}

    uniq := GetOutputSlice(slice)
    assert.Equal(len(uniq), 5)
    assert.Equal(uniq[0], "2 Pr M Father")
    assert.Equal(uniq[1], "1 FR O Mother")
    assert.Equal(uniq[2], "1 NR M Son")
    assert.Equal(uniq[3], "1 DR T Dad")
    assert.Equal(uniq[4], "2 XR y SON")
    

    numFields = 0
    numChars = 0
    caseInsensitive = false
    numOccurrencesStrings = false
}
