package closest_pair

import (
	"math"
)

func divideAndConquer(P []point) *pair {
	n := len(P)

	if n == 2 {
		return makePair(P[0], P[1])
	}

	xP := make([]point, n)
	yP := make([]point, n)

	for i := 0; i < n; i++ {
		xP[i] = P[i]
		yP[i] = P[i]
	}

	_, pair := _divideAndConquer(xP, yP)

	return pair
}

func _divideAndConquer(xP, yP []point) (float64, *pair) {
	n := len(xP)

	if n <= 3 {
		p := bruteForce(xP)
		return p.distance, p
	}

	xL := xP[:n>>1]
	xR := xP[n>>1:]

	var yL, yR []point
	xMiddle := xL[0].x

	// yL ←  { p ∈ yP : px ≤ xMiddle }
	// yR ←  { p ∈ yP : px > xMiddle }
	for _, p := range yP {
		if p.x <= xMiddle {
			yL = append(yL, p)
		} else {
			yR = append(yL, p)
		}
	}

	// (dL, pairL) ←  closestPair of (xL, yL)
	// (dR, pairR) ←  closestPair of (xR, yR)
	dL, pairL := _divideAndConquer(xL, yL)
	dR, pairR := _divideAndConquer(xR, yR)

	dMin, pairMin := dR, pairR

	if dL < dR {
		dMin, pairMin = dL, pairL
	}

	var yS []point

	// yS ← { p ∈ yP : |xMiddle - px| < dMin }
	for i := 0; i < len(yP); i++ {
		if math.Abs(yP[i].x-xMiddle) < dMin {
			yS = append(yS, yP[i])
		}
	}

	nS := len(yS)

	if nS > 1 {
		closestPair := pairMin
		for i := 1; i < nS-1; i++ {
			k := i + 1
			for k <= nS && math.Abs(yS[k].y-yS[i].y) < dMin {
				tempPair := makePair(yS[k], yS[i])
				if tempPair.distance < closestPair.distance {
					closestPair = tempPair
				}
				k++
			}
		}
		return closestPair.distance, closestPair
	} else {
		return dMin, pairMin
	}
}
