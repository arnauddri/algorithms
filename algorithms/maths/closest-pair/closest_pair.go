package closest_pair

import (
	"math"
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

type pair struct {
	point1   point
	point2   point
	distance float64
}

type Pair interface {
	calcDistance() float64
	toString() string
}

func makePair(A point, B point) *pair {
	P := new(pair)
	P.point1 = A
	P.point2 = B
	P.distance = calcDistance(A, B)

	return P
}

func calcDistance(A point, B point) float64 {
	xdist := A.x - B.x
	ydist := A.y - B.y

	return math.Sqrt(float64(xdist*xdist + ydist*ydist))
}

func (P *pair) toString() string {
	return P.point1.toString() + "-" + P.point2.toString() + "-" + strconv.FormatFloat(P.distance, 'f', -1, 64)
}

func DivideAndConquer(P []point) *pair {
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

	_, pair := divideAndConquer(xP, yP)

	return pair
}

func divideAndConquer(xP, yP []point) (float64, *pair) {
	n := len(xP)

	if n <= 3 {
		p := BruteForce(xP)
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
	dL, pairL := divideAndConquer(xL, yL)
	dR, pairR := divideAndConquer(xR, yR)

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

func BruteForce(P []point) *pair {
	n := len(P)

	if n < 2 {
		return nil
	}

	minPair := makePair(P[0], P[1])
	tempPair := makePair(P[0], P[1])

	min := minPair.distance
	for i := 0; i < n-1; i++ {
		tempPair.point1 = P[i]
		for j := 0; j < n; j++ {
			tempPair.point2 = P[j]
			if min > tempPair.distance {
				minPair = tempPair
				min = tempPair.distance
			}
		}
	}

	return minPair
}
