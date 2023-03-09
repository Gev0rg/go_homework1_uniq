package uniq

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var forGetUniqSlice = []struct {
	input   []string
	compare []string
	result  []string
}{
	{
		[]string{"Father", "Father", "Mother", "Son", "Son", "Son"},
		[]string{"Father", "Father", "Mother", "Son", "Son", "Son"},
		[]string{"Father", "Mother", "Son"},
	},
}

func TestGetUniqSlice(t *testing.T) {
	for _, test := range forGetUniqSlice {
		uniq, _ := GetUniqSlice(test.input, test.compare)
		require.Equal(t, test.result, uniq)
	}
}

var forGetOutputSlice = []struct {
	input   []string
	result  []string
	options Options
}{
	{
		[]string{"Father", "Father", "Mother", "Son", "Son", "Son"},
		[]string{"Father", "Mother", "Son"},
		Options{},
	},
	{
		[]string{"My Father", "Your Father", "Our Mother", "My Son", "Your Son", "Our Son"},
		[]string{"2 My Father", "3 My Son"},
		Options{true, true, false, 1, 0, false},
	},
	{
		[]string{"M Father", "Y father", "O Mother", "M Son", "T Dad", "y SON", "O sOn"},
		[]string{"O Mother", "M Son", "T Dad"},
		Options{false, false, true, 0, 1, true},
	},
	{
		[]string{"Pr M Father", "MR Y father", "FR O Mother", "NR M Son", "DR T Dad", "XR y SON", "Sr O sOn"},
		[]string{"2 Pr M Father", "1 FR O Mother", "1 NR M Son", "1 DR T Dad", "2 XR y SON"},
		Options{true, false, false, 1, 2, true},
	},
}

func TestGetOutputSlice1(t *testing.T) {
	for _, test := range forGetOutputSlice {
		uniq := GetOutputSlice(test.input, test.options)
		require.Equal(t, test.result, uniq)
	}
}
