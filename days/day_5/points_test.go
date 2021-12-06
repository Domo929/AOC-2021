package day_5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Line_PointsVOrH(t *testing.T) {
	t.Run("same point", func(t *testing.T) {
		line := line{
			A: point{
				X: 1,
				Y: 1,
			},
			B: point{
				X: 1,
				Y: 1,
			},
		}
		expectedPoints := []point{
			{X: 1, Y: 1},
		}
		assert.Equal(t, expectedPoints, line.pointsVorH())
	})
	t.Run("horizontal line to the right", func(t *testing.T) {
		line := line{
			A: point{
				X: 0,
				Y: 5,
			},
			B: point{
				X: 3,
				Y: 5,
			},
		}
		expectedPoints := []point{
			{X: 0, Y: 5},
			{X: 1, Y: 5},
			{X: 2, Y: 5},
			{X: 3, Y: 5},
		}
		assert.Equal(t, expectedPoints, line.pointsVorH())
	})
	t.Run("horizontal line to the left", func(t *testing.T) {
		line := line{
			A: point{
				X: 3,
				Y: 5,
			},
			B: point{
				X: 0,
				Y: 5,
			},
		}
		expectedPoints := []point{
			{X: 3, Y: 5},
			{X: 2, Y: 5},
			{X: 1, Y: 5},
			{X: 0, Y: 5},
		}
		assert.Equal(t, expectedPoints, line.pointsVorH())
	})

	t.Run("vertical line up", func(t *testing.T) {
		line := line{
			A: point{
				X: 3,
				Y: 0,
			},
			B: point{
				X: 3,
				Y: 2,
			},
		}
		expectedPoints := []point{
			{X: 3, Y: 0},
			{X: 3, Y: 1},
			{X: 3, Y: 2},
		}
		assert.Equal(t, expectedPoints, line.pointsVorH())
	})

	t.Run("vertical line down", func(t *testing.T) {
		line := line{
			A: point{
				X: 3,
				Y: 2,
			},
			B: point{
				X: 3,
				Y: 0,
			},
		}
		expectedPoints := []point{
			{X: 3, Y: 2},
			{X: 3, Y: 1},
			{X: 3, Y: 0},
		}
		assert.Equal(t, expectedPoints, line.pointsVorH())
	})
}

func Test_line_PointsDiag(t *testing.T) {
	t.Run("up right", func(t *testing.T) {
		line := line{
			A: point{
				X: 2,
				Y: 2,
			},
			B: point{
				X: 4,
				Y: 4,
			},
		}
		expectedPoints := []point{
			{X: 2, Y: 2},
			{X: 3, Y: 3},
			{X: 4, Y: 4},
		}
		assert.Equal(t, expectedPoints, line.pointsDiag())
	})

	t.Run("down right", func(t *testing.T) {
		line := line{
			A: point{
				X: 2,
				Y: 3,
			},
			B: point{
				X: 4,
				Y: 1,
			},
		}
		expectedPoints := []point{
			{X: 2, Y: 3},
			{X: 3, Y: 2},
			{X: 4, Y: 1},
		}
		assert.Equal(t, expectedPoints, line.pointsDiag())
	})

	t.Run("up left", func(t *testing.T) {
		line := line{
			A: point{
				X: 4,
				Y: 1,
			},
			B: point{
				X: 2,
				Y: 3,
			},
		}
		expectedPoints := []point{
			{X: 4, Y: 1},
			{X: 3, Y: 2},
			{X: 2, Y: 3},
		}
		assert.Equal(t, expectedPoints, line.pointsDiag())
	})
	t.Run("down left", func(t *testing.T) {
		line := line{
			A: point{
				X: 4,
				Y: 4,
			},
			B: point{
				X: 2,
				Y: 2,
			},
		}
		expectedPoints := []point{
			{X: 4, Y: 4},
			{X: 3, Y: 3},
			{X: 2, Y: 2},
		}
		assert.Equal(t, expectedPoints, line.pointsDiag())
	})

}
