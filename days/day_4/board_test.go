package day_4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isBingo(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		b := board{squares: [25]square{}}
		assert.False(t, b.isBingo())
	})

	t.Run("random squares", func(t *testing.T) {
		b := board{squares: [25]square{}}
		b.squares[23].marked = true
		b.squares[11].marked = true
		b.squares[4].marked = true
		b.squares[2].marked = true
		assert.False(t, b.isBingo())
	})

	t.Run("row", func(t *testing.T) {
		b := board{squares: [25]square{}}
		for i := 0; i < 5; i++ {
			b.squares[i].marked = true
		}
		assert.True(t, b.isBingo())
	})

	t.Run("col", func(t *testing.T) {
		b := board{squares: [25]square{}}
		for i := 0; i < 25; i += 5 {
			b.squares[i].marked = true
		}
		assert.True(t, b.isBingo())
	})

	//t.Run("diag back slash", func(t *testing.T) {
	//	b := board{squares: [25]square{}}
	//	for i := 0; i < 25; i += 6 {
	//		b.squares[i].marked = true
	//	}
	//	assert.True(t, b.isBingo())
	//})
	//
	//t.Run("diag forward slash", func(t *testing.T) {
	//	b := board{squares: [25]square{}}
	//	for i := 4; i < 25; i += 4 {
	//		b.squares[i].marked = true
	//	}
	//	assert.True(t, b.isBingo())
	//})

}
