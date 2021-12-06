package day_5

import (
	"math"
)

type point struct {
	X, Y int
}

type line struct {
	A, B point
}

func (l line) pointsDiag() []point {
	points := make([]point, 0)
	xIncrement := l.A.X <= l.B.X
	yIncrement := l.A.Y <= l.B.Y

	diff := int(math.Abs(float64(l.A.X - l.B.X)))
	x, y := l.A.X, l.A.Y
	for i := 0; i < diff; i++ {
		points = append(points, point{X: x, Y: y})
		x = crement(xIncrement, x)
		y = crement(yIncrement, y)
	}
	points = append(points, point{X: x, Y: y})

	return points
}

func crement(isIncrement bool, val int) int {
	if isIncrement {
		return val + 1
	}
	return val - 1
}

func (l line) isDiag() bool {
	xDiff := int(math.Abs(float64(l.A.X - l.B.X)))
	yDiff := int(math.Abs(float64(l.A.Y - l.B.Y)))
	return xDiff == yDiff
}

func (l line) pointsVorH() []point {
	points := make([]point, 0)
	if l.A.X == l.B.X {
		if l.A.Y < l.B.Y {
			for y := l.A.Y; y <= l.B.Y; y++ {
				points = append(points, point{
					X: l.A.X,
					Y: y,
				})
			}
		} else {
			for y := l.A.Y; y >= l.B.Y; y-- {
				points = append(points, point{
					X: l.A.X,
					Y: y,
				})
			}
		}
	} else { // ys constant
		if l.A.X < l.B.X {
			for x := l.A.X; x <= l.B.X; x++ {
				points = append(points, point{
					X: x,
					Y: l.A.Y,
				})
			}
		} else {
			for x := l.A.X; x >= l.B.X; x-- {
				points = append(points, point{
					X: x,
					Y: l.A.Y,
				})
			}
		}
	}
	return points
}

func (l line) isFlat() bool {
	return l.A.X == l.B.X || l.A.Y == l.B.Y
}
