package day_8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeDigits(t *testing.T) {
	digits := []string{
		"acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab",
	}
	expectedNumbers := map[string]int{
		"acedgfb": 8, "cdfbe": 5, "gcdfa": 2, "fbcad": 3, "dab": 7, "cefabd": 9, "cdfgeb": 6, "eafb": 4, "cagedb": 0, "ab": 1,
	}

	actualNumbers := decodeDigits(digits)
	assert.Equal(t, expectedNumbers, actualNumbers)
}
