package day_7

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSumMap(t *testing.T) {
	sm := make(sumMap)

	t.Run("repeat", func(t *testing.T) {
		fuel := sm.sumFuel(2)
		require.Equal(t, 3, fuel)

		fuel = sm.sumFuel(2)
		assert.Equal(t, 3, fuel)
	})

	t.Run("0", func(t *testing.T) {
		fuel := sm.sumFuel(0)
		assert.Equal(t, 0, fuel)
	})

	t.Run("1", func(t *testing.T) {
		fuel := sm.sumFuel(1)
		assert.Equal(t, 1, fuel)
	})

	t.Run("2", func(t *testing.T) {
		fuel := sm.sumFuel(2)
		assert.Equal(t, 3, fuel)
	})
}
