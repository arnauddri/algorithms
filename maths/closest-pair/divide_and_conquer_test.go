package closest_pair

import (
	"fmt"
	"testing"
)

func TestDivideAndConquer(t *testing.T) {
	Points := make([]point, 10)
	for i := 0; i < 10; i++ {
		a := makePoint(float64(i), float64(i*i))

		Points[i] = *a
	}

	A := makePoint(float64(0), float64(0))
	B := makePoint(float64(1), float64(1))
	P := divideAndConquer(Points)
	if P.point1 != *A || P.point2 != *B {
		fmt.Println(Points)
		fmt.Println(P)
		t.Error()
	}
}
