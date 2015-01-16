package closest_pair

import (
	"strconv"
)

type point struct {
	x float64
	y float64
}

type Point interface {
	toString() string
}

func makePoint(x float64, y float64) *point {
	A := new(point)
	A.x = x
	A.y = y

	return A
}

func (A *point) toString() string {
	return "(" + strconv.FormatFloat(A.x, 'f', -1, 64) + "," + strconv.FormatFloat(A.y, 'f', -1, 64) + ")"
}
