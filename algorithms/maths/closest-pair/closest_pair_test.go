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
	P := DivideAndConquer(Points)
	if P.point1 != *A || P.point2 != *B {
		fmt.Println(Points)
		fmt.Println(P)
		t.Error()
	}
}
func TestBruteForce(t *testing.T) {
	Points := make([]point, 10)
	for i := 0; i < 10; i++ {
		a := makePoint(float64(i), float64(i*i))

		Points[i] = *a
	}

	A := makePoint(float64(0), float64(0))
	B := makePoint(float64(1), float64(1))
	P := BruteForce(Points)
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
		BruteForce(Points)
	}
}

func benchmarkDivideAndConquer(n int, b *testing.B) {
	Points := make([]point, n)
	for i := 0; i < n; i++ {
		a := makePoint(float64(i), float64(i*i))

		Points[i] = *a
	}

	for i := 0; i < b.N; i++ {
		DivideAndConquer(Points)
	}
}

//func BenchmarkDivideAndConquer100(b *testing.B)      { benchmarkDivideAndConquer(100, b) }
//func BenchmarkDivideAndConquer1000(b *testing.B)     { benchmarkDivideAndConquer(1000, b) }
//func BenchmarkDivideAndConquer10000(b *testing.B)    { benchmarkDivideAndConquer(10000, b) }
//func BenchmarkDivideAndConquer100000(b *testing.B)   { benchmarkDivideAndConquer(100000, b) }
//func BenchmarkDivideAndConquer1000000(b *testing.B)  { benchmarkDivideAndConquer(1000000, b) }
//func BenchmarkDivideAndConquer10000000(b *testing.B) { benchmarkDivideAndConquer(10000000, b) }

func BenchmarkDivideAndConquer(b *testing.B) {
	for i := 0; i < 10; i++ {

		benchmarkDivideAndConquer(100, b)
	}
}
