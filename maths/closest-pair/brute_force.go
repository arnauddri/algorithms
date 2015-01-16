package closest_pair

import ()

func bruteForce(P []point) *pair {
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
