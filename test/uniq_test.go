package main

import (
    "github.com/Gev0rg/go_homework1_uniq/uniq"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUniqSlice(t *testing.T) {
	assert := assert.New(t)
    slice := []string{"Father", "Father", "Mother", "Son", "Son", "Son"}
    uniq := GetUniqSlice(slice)
    assert.Equal(len(uniq), 3)
    assert.Equal(uniq[1], "Mother")
}
