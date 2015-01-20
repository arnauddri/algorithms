package closest_pair

import (
	"fmt"
	"testing"
)

func TestBruteForce(t *testing.T) {
	Points := make([]point, 10)
	for i := 0; i < 10; i++ {
		a := makePoint(float64(i), float64(i*i))

		Points[i] = *a
	}

	A := makePoint(float64(0), float64(0))
	B := makePoint(float64(1), float64(1))
	P := bruteForce(Points)
	if P.point1 != *A || P.point2 != *B {
		fmt.Println(P.point1, *A, *B)
		t.Error()
	}
}

func BenchmarkBruteForce(b *testing.B) {
	Points := make([]point, 10)
	for i := 0; i < 10; i++ {
		a := makePoint(float64(i), float64(i*i))

		Points[i] = *a
	}

	for i := 0; i < b.N; i++ {
		bruteForce(Points)
	}
}
