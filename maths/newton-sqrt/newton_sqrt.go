package newton_sqrt

import (
	"math"
)

func newton_sqrt(n float64, precision, maxIterations float64) float64 {
	upperBound := float64(n)
	lowerBound := float64(0)

	var square float64
	var i int
	var x float64

	for math.Abs(square-n) > precision && float64(i) < maxIterations {
		i++
		x = (upperBound-lowerBound)/2 + lowerBound
		square = x * x

		if square < n {
			lowerBound = x
		} else {
			upperBound = x
		}
	}

	floorX := math.Floor(x)

	if floorX*floorX == n {
		x = floorX
	} else if (floorX+1)*(floorX+1) == n {
		x = floorX + 1
	}

	return x
}
